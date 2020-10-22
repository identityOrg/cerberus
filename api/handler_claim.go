package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetClaims(ctx echo.Context, params GetClaimsParams) error {
	pageNumber := 0
	pageSize := 5
	if params.PageNumber != nil {
		pageNumber = *(params.PageNumber)
	}
	if params.PageSize != nil {
		pageSize = *(params.PageSize)
	}
	claims, total, err := c.ScopeClaimStore.GetAllClaims(ctx.Request().Context(), uint(pageNumber), uint(pageSize))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	} else {
		page := &ClaimPage{
			Page: Page{
				PageNumber: pageNumber,
			},
		}
		page.PageTotal = CalculatePageCount(total, uint(pageSize))
		for _, claim := range claims {
			cl := Claim{
				Description: claim.Description,
				Id:          int(claim.ID),
				Name:        claim.Name,
			}
			page.Claims = append(page.Claims, cl)
		}
		return ctx.JSON(http.StatusOK, page)
	}
}

func (c *CerberusAPI) CreateClaim(ctx echo.Context) error {
	createReq := &Claim{}
	err := ctx.Bind(createReq)
	if err != nil {
		return err
	}
	_, err = c.ScopeClaimStore.CreateClaim(ctx.Request().Context(), createReq.Name, createReq.Description)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusCreated)
}

func (c *CerberusAPI) DeleteClaim(ctx echo.Context, id int) error {
	err := c.ScopeClaimStore.DeleteClaim(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c *CerberusAPI) GetClaim(ctx echo.Context, id int) error {
	claim, err := c.ScopeClaimStore.GetClaim(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	cl := &Claim{
		Description: claim.Description,
		Id:          int(claim.ID),
		Name:        claim.Name,
	}
	return ctx.JSON(http.StatusOK, cl)
}

func (c *CerberusAPI) UpdateClaim(ctx echo.Context, id int) error {
	cl := &Claim{}
	err := ctx.Bind(cl)
	if err != nil {
		return err
	}
	err = c.ScopeClaimStore.UpdateClaim(ctx.Request().Context(), uint(id), cl.Description)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) FindClaimByName(ctx echo.Context, params FindClaimByNameParams) error {
	claim, err := c.ScopeClaimStore.FindClaimByName(ctx.Request().Context(), params.Name)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	cl := &Claim{
		Description: claim.Description,
		Id:          int(claim.ID),
		Name:        claim.Name,
	}
	return ctx.JSON(http.StatusOK, cl)
}
