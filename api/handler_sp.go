package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetServiceProviders(ctx echo.Context, params GetServiceProvidersParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) CreateServiceProvider(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) DeleteServiceProvider(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GetServiceProvider(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) UpdateServiceProvider(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GetCredentials(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GenerateCredentials(ctx echo.Context, id int, params GenerateCredentialsParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) FindServiceProvider(ctx echo.Context, params FindServiceProviderParams) error {
	panic("implement me")
}

func (c *CerberusAPI) PatchServiceProvider(ctx echo.Context, id int) error {
	panic("implement me")
}

func (c *CerberusAPI) DeactivateServiceProvider(ctx echo.Context, id int) error {
	panic("implement me")
}

func (c *CerberusAPI) ActivateServiceProvider(ctx echo.Context, id int) error {
	panic("implement me")
}
