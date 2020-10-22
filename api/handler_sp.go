package api

import (
	"github.com/identityOrg/cerberus/core/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	DefaultSignatureAlg = "RS256"
)

func (c *CerberusAPI) GetServiceProviders(ctx echo.Context, params GetServiceProvidersParams) error {
	var page uint = 0
	var size uint = 5
	if params.PageNumber != nil {
		page = uint(*params.PageNumber)
	}
	if params.PageSize != nil {
		size = uint(*params.PageSize)
	}
	sps, total, err := c.SPStoreService.FindAllSP(ctx.Request().Context(), page, size)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	spp := ServiceProviderSummaryPage{
		Page: Page{
			PageNumber: int(page),
		},
	}
	spp.PageTotal = CalculatePageCount(total, size)
	for _, sp := range sps {
		spSumm := ServiceProviderSummary{
			Description: sp.Description,
			Id:          int(sp.ID),
			Name:        sp.Name,
		}
		spp.ServiceProviders = append(spp.ServiceProviders, spSumm)
	}
	return ctx.JSON(http.StatusOK, spp)
}

func (c *CerberusAPI) CreateServiceProvider(ctx echo.Context) error {
	sp := &ServiceProvider{}
	err := ctx.Bind(sp)
	if err != nil {
		return err
	}
	metadata := c.convertToMetadata(sp)

	_, err = c.SPStoreService.CreateSP(ctx.Request().Context(), sp.Name, sp.Description, metadata)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusCreated)
}

func (c *CerberusAPI) DeleteServiceProvider(ctx echo.Context, id int) error {
	err := c.SPStoreService.DeleteSP(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c *CerberusAPI) GetServiceProvider(ctx echo.Context, id int) error {
	provider, err := c.SPStoreService.GetSP(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	sp := c.convertToSP(provider)
	return ctx.JSON(http.StatusOK, sp)
}

func (c *CerberusAPI) UpdateServiceProvider(ctx echo.Context, id int) error {
	apiSp := &ServiceProvider{}
	err := ctx.Bind(apiSp)
	if err != nil {
		return err
	}
	metadata := c.convertToMetadata(apiSp)
	err = c.SPStoreService.UpdateSP(ctx.Request().Context(), uint(id), apiSp.Public, metadata)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) GetCredentials(ctx echo.Context, id int) error {
	sp, err := c.SPStoreService.GetSP(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	spCred := &ServiceProviderCredentials{
		ClientId:     sp.ClientID,
		ClientSecret: sp.ClientSecret,
	}
	return ctx.JSON(http.StatusOK, spCred)
}

func (c *CerberusAPI) GenerateCredentials(ctx echo.Context, id int) error {
	genCred := &RegenerateCredentials{}
	err := ctx.Bind(genCred)
	if err != nil {
		return err
	}
	credentials, secret, err := c.SPStoreService.ResetClientCredentials(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	spCred := &ServiceProviderCredentials{
		ClientId:     credentials,
		ClientSecret: secret,
	}
	return ctx.JSON(http.StatusOK, spCred)
}

func (c *CerberusAPI) FindServiceProvider(ctx echo.Context, params FindServiceProviderParams) error {
	if params.ClientId != nil {
		sp, err := c.SPStoreService.FindSPByClientId(ctx.Request().Context(), *(params.ClientId))
		if err != nil {
			return &ApiError{
				ErrorCode: "error",
				Message:   err.Error(),
			}
		}
		apiSp := c.convertToSP(sp)
		return ctx.JSON(http.StatusOK, apiSp)
	} else if params.Name != nil {
		sp, err := c.SPStoreService.FindSPByName(ctx.Request().Context(), *(params.Name))
		if err != nil {
			return &ApiError{
				ErrorCode: "error",
				Message:   err.Error(),
			}
		}
		apiSp := c.convertToSP(sp)
		return ctx.JSON(http.StatusOK, apiSp)
	} else {
		return &ApiError{
			ErrorCode: "error",
			Message:   "either client_id or name parameter is required",
		}
	}
}

func (c *CerberusAPI) PatchServiceProvider(ctx echo.Context, _ int) error {
	return ctx.String(http.StatusNotImplemented, "implement me")
}

func (c *CerberusAPI) UpdateServiceProviderStatus(ctx echo.Context, id int) error {
	update := &StatusUpdate{}
	err := ctx.Bind(update)
	if err != nil {
		return err
	}
	if update.Active {
		err = c.SPStoreService.ActivateSP(ctx.Request().Context(), uint(id))
	} else {
		err = c.SPStoreService.DeactivateSP(ctx.Request().Context(), uint(id))
	}
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) convertToSP(sp *models.ServiceProviderModel) *ServiceProvider {
	apiSp := &ServiceProvider{}
	apiSp.Scope = sp.Metadata.Scopes
	apiSp.RedirectUris = sp.Metadata.RedirectUris
	apiSp.GrantTypes = sp.Metadata.GrantTypes
	apiSp.ApplicationType = sp.Metadata.ApplicationType
	apiSp.Contacts = sp.Metadata.Contacts
	apiSp.DefaultAcrValues = sp.Metadata.DefaultAcrValues
	apiSp.DefaultMaxAge = sp.Metadata.DefaultMaxAge
	apiSp.IdTokenEncryptedResponseAlg = sp.Metadata.IdTokenEncryptedResponseAlg
	apiSp.IdTokenEncryptedResponseEnc = sp.Metadata.IdTokenEncryptedResponseEnc
	if sp.Metadata.IdTokenSignedResponseAlg != "" {
		apiSp.IdTokenSignedResponseAlg = &sp.Metadata.IdTokenSignedResponseAlg
	}
	apiSp.RequestObjectEncryptionAlg = sp.Metadata.RequestObjectEncryptionAlg
	apiSp.RequestObjectEncryptionEnc = sp.Metadata.RequestObjectEncryptionEnc
	apiSp.RequestObjectSigningAlg = sp.Metadata.RequestObjectSigningAlg
	apiSp.UserinfoEncryptedResponseAlg = sp.Metadata.UserinfoEncryptedResponseAlg
	apiSp.UserinfoEncryptedResponseEnc = sp.Metadata.UserinfoEncryptedResponseEnc
	apiSp.UserinfoSignedResponseAlg = sp.Metadata.UserinfoSignedResponseAlg
	apiSp.InitiateLoginUri = sp.Metadata.InitiateLoginUri
	apiSp.JwksUri = sp.Metadata.JwksUri
	apiSp.LogoUri = sp.Metadata.LogoUri
	apiSp.ClientUri = sp.Metadata.ClientUri
	apiSp.PolicyUri = sp.Metadata.PolicyUri
	apiSp.RequireAuthTime = sp.Metadata.RequireAuthTime
	//apiSp.SectorIdentifierUri = sp.Sec
	//apiSp.SubjectType = sp.Sub
	if sp.Metadata.TokenEndpointAuthMethod != "" {
		apiSp.TokenEndpointAuthMethod = &sp.Metadata.TokenEndpointAuthMethod
	}
	apiSp.TokenEndpointAuthSigningAlg = sp.Metadata.TokenEndpointAuthSigningAlg

	apiSp.Name = sp.Name
	apiSp.Description = sp.Description
	return apiSp
}

func (c *CerberusAPI) convertToMetadata(sp *ServiceProvider) *models.ServiceProviderMetadata {
	metadata := &models.ServiceProviderMetadata{}
	metadata.Scopes = sp.Scope
	metadata.RedirectUris = sp.RedirectUris
	metadata.GrantTypes = sp.GrantTypes
	metadata.ApplicationType = sp.ApplicationType
	metadata.Contacts = sp.Contacts
	metadata.DefaultAcrValues = sp.DefaultAcrValues
	metadata.DefaultMaxAge = sp.DefaultMaxAge
	metadata.IdTokenEncryptedResponseAlg = sp.IdTokenEncryptedResponseAlg
	metadata.IdTokenEncryptedResponseEnc = sp.IdTokenEncryptedResponseEnc
	if sp.IdTokenSignedResponseAlg != nil {
		metadata.IdTokenSignedResponseAlg = *sp.IdTokenSignedResponseAlg
	} else {
		metadata.IdTokenSignedResponseAlg = DefaultSignatureAlg
	}
	metadata.RequestObjectEncryptionAlg = sp.RequestObjectEncryptionAlg
	metadata.RequestObjectEncryptionEnc = sp.RequestObjectEncryptionEnc
	metadata.RequestObjectSigningAlg = sp.RequestObjectSigningAlg
	metadata.UserinfoEncryptedResponseAlg = sp.UserinfoEncryptedResponseAlg
	metadata.UserinfoEncryptedResponseEnc = sp.UserinfoEncryptedResponseEnc
	metadata.UserinfoSignedResponseAlg = sp.UserinfoSignedResponseAlg
	metadata.InitiateLoginUri = sp.InitiateLoginUri
	metadata.JwksUri = sp.JwksUri
	metadata.LogoUri = sp.LogoUri
	metadata.ClientUri = sp.ClientUri
	metadata.PolicyUri = sp.PolicyUri
	metadata.RequireAuthTime = sp.RequireAuthTime
	//metadata.SectorIdentifierUri = sp.Sec
	//metadata.SubjectType = sp.Sub
	if sp.TokenEndpointAuthMethod != nil {
		metadata.TokenEndpointAuthMethod = *sp.TokenEndpointAuthMethod
	} else {
		metadata.TokenEndpointAuthMethod = "client_secret_basic"
	}
	metadata.TokenEndpointAuthSigningAlg = sp.TokenEndpointAuthSigningAlg
	return metadata
}
