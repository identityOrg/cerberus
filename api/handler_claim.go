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
		e := &Error{
			ErrorCode: convertStringP("error"),
			Message:   convertStringP(err.Error()),
		}
		return e
	} else {
		page := &ClaimPage{
			Page: Page{
				PageNumber: pageNumber,
				PageTotal:  int(total),
			},
		}
		for _, claim := range claims {
			cl := Claim{
				Description: &claim.Description,
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
	desc := ""
	if createReq.Description != nil {
		desc = *createReq.Description
	}
	_, err = c.ScopeClaimStore.CreateClaim(ctx.Request().Context(), createReq.Name, desc)
	if err != nil {
		return &Error{
			ErrorCode: convertStringP("error"),
			Message:   convertStringP(err.Error()),
		}
	}
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) DeleteClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) GetClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) UpdateClaim(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) FindClaimByName(ctx echo.Context, params FindClaimByNameParams) error {
	panic("implement me")
}
