package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetScopes(ctx echo.Context, params GetScopesParams) error {
	var page uint = 0
	var size uint = 5
	if params.PageNumber != nil {
		page = uint(*params.PageNumber)
	}
	if params.PageSize != nil {
		size = uint(*params.PageSize)
	}
	scopes, total, err := c.ScopeClaimStore.GetAllScopes(ctx.Request().Context(), page, size)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	scopePage := &ScopePage{
		Page: Page{
			PageNumber: int(page),
		},
	}
	scopePage.PageTotal = CalculatePageCount(total, size)
	for _, scope := range scopes {
		cl := Scope{
			Description: scope.Description,
			Id:          int(scope.ID),
			Name:        scope.Name,
		}
		scopePage.Scopes = append(scopePage.Scopes, cl)
	}
	return ctx.JSON(http.StatusOK, scopePage)
}

func (c *CerberusAPI) CreateScope(ctx echo.Context) error {
	scope := &Scope{}
	err := ctx.Bind(scope)
	if err != nil {
		return err
	}
	_, err = c.ScopeClaimStore.CreateScope(ctx.Request().Context(), scope.Name, scope.Description)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusCreated)
}

func (c *CerberusAPI) DeleteScope(ctx echo.Context, id int) error {
	err := c.ScopeClaimStore.DeleteScope(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c *CerberusAPI) GetScope(ctx echo.Context, id int) error {
	scope, err := c.ScopeClaimStore.GetScope(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	sc := &Scope{
		Description: scope.Description,
		Id:          int(scope.ID),
		Name:        scope.Name,
	}
	return ctx.JSON(http.StatusOK, sc)
}

func (c *CerberusAPI) UpdateScope(ctx echo.Context, id int) error {
	scope := &Scope{}
	err := ctx.Bind(scope)
	if err != nil {
		return err
	}
	err = c.ScopeClaimStore.UpdateScope(ctx.Request().Context(), uint(id), scope.Description)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) FindScopeByName(ctx echo.Context, params FindScopeByNameParams) error {
	scope, err := c.ScopeClaimStore.FindScopeByName(ctx.Request().Context(), params.Name)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	sc := &Scope{
		Description: scope.Description,
		Id:          int(scope.ID),
		Name:        scope.Name,
	}
	return ctx.JSON(http.StatusOK, sc)
}

func (c *CerberusAPI) RemoveClaimFromScope(ctx echo.Context, id int, claimId int) error {
	err := c.ScopeClaimStore.RemoveClaimFromScope(ctx.Request().Context(), uint(id), uint(claimId))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) AddClaimToScope(ctx echo.Context, id int) error {
	addClaimToScope := &AddClaimToScope{}
	err := ctx.Bind(addClaimToScope)
	if err != nil {
		return err
	}
	err = c.ScopeClaimStore.AddClaimToScope(ctx.Request().Context(), uint(id), uint(addClaimToScope.ClaimId))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}
