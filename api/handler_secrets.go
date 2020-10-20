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
		return &Error{
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
	panic("implement me")
}

func (c *CerberusAPI) FindSecretChannelByAlgouse(ctx echo.Context, params FindSecretChannelByAlgouseParams) error {
	panic("implement me")
}

func (c *CerberusAPI) FindSecretChannelByName(ctx echo.Context, params FindSecretChannelByNameParams) error {
	panic("implement me")
}

func (c *CerberusAPI) DeleteSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}

func (c *CerberusAPI) GetSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}

func (c *CerberusAPI) RenewSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}

func (c *CerberusAPI) UpdateSecretChannel(ctx echo.Context, id string) error {
	panic("implement me")
}
