package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetSecretChannels(ctx echo.Context, params GetSecretChannelsParams) error {
	var page uint = 0
	var size uint = 5
	if params.PageNumber != nil {
		page = uint(*params.PageNumber)
	}
	if params.PageSize != nil {
		size = uint(*params.PageSize)
	}
	channels, total, err := c.SecretStoreStore.GetAllChannels(ctx.Request().Context(), page, size)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	scPage := &SecretChannelPage{
		Page: Page{
			PageNumber: int(page),
		},
	}
	scPage.PageTotal = CalculatePageCount(total, size)
	for _, ch := range channels {
		cl := SecretChannelSummary{
			Algorithm: ch.Algorithm,
			Id:        int(ch.ID),
			KeyUsage:  ch.Use,
			Name:      ch.Name,
		}
		scPage.Channels = append(scPage.Channels, cl)
	}
	return ctx.JSON(http.StatusOK, scPage)
}

func (c *CerberusAPI) CreateSecretChannel(ctx echo.Context) error {
	sc := &SecretChannel{}
	err := ctx.Bind(sc)
	if err != nil {
		return err
	}
	_, err = c.SecretStoreStore.CreateChannel(ctx.Request().Context(), sc.Name, sc.Algorithm, sc.KeyUsage, uint(sc.ValidityDay))
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (c *CerberusAPI) FindSecretChannelByAlgouse(ctx echo.Context, params FindSecretChannelByAlgouseParams) error {
	ch, err := c.SecretStoreStore.GetChannelByAlgoUse(ctx.Request().Context(), params.Algo, params.Use)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, ch)
}

func (c *CerberusAPI) FindSecretChannelByName(ctx echo.Context, params FindSecretChannelByNameParams) error {
	ch, err := c.SecretStoreStore.GetChannelByName(ctx.Request().Context(), params.Name)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, ch)
}

func (c *CerberusAPI) DeleteSecretChannel(ctx echo.Context, id int) error {
	err := c.SecretStoreStore.DeleteChannel(ctx.Request().Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c *CerberusAPI) GetSecretChannel(ctx echo.Context, id int) error {
	ch, err := c.SecretStoreStore.GetChannel(ctx.Request().Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, ch)
}

func (c *CerberusAPI) RenewSecretChannel(ctx echo.Context, id int) error {
	err := c.SecretStoreStore.RenewSecret(ctx.Request().Context(), uint(id))
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
