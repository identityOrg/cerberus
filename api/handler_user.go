package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetUsers(ctx echo.Context, params GetUsersParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) CreateUser(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) InitiatePasswordRecovery(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) ResetUserPassword(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) DeleteUser(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GetUser(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) UpdateUser(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) ChangeUserPassword(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) UpdateUserStatus(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) FindUser(ctx echo.Context, params FindUserParams) error {
	panic("implement me")
}
