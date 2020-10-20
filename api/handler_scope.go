package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetScopes(ctx echo.Context, params GetScopesParams) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) CreateScope(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) DeleteScope(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GetScope(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) UpdateScope(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) FindScopeByName(ctx echo.Context, params FindScopeByNameParams) error {
	panic("implement me")
}

func (c *CerberusAPI) RemoveClaimFromScope(ctx echo.Context, id int, claimId int) error {
	panic("implement me")
}

func (c *CerberusAPI) AddClaimToScope(ctx echo.Context, id int, claimId int) error {
	panic("implement me")
}
