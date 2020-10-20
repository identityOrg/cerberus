package api

import (
	"github.com/labstack/echo/v4"
)

func (c *CerberusAPI) GetSecretChannels(ctx echo.Context, params GetSecretChannelsParams) error {
	panic("implement me")
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
