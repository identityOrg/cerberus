package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *CerberusAPI) GetUsers(ctx echo.Context, params GetUsersParams) error {
	var page uint = 0
	var size uint = 5
	if params.PageNumber != nil {
		page = uint(*params.PageNumber)
	}
	if params.PageSize != nil {
		size = uint(*params.PageSize)
	}
	users, total, err := c.UserStoreService.FindAllUser(ctx.Request().Context(), page, size)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	up := &UserSummaryPage{
		Page: Page{
			PageNumber: int(page),
		},
	}
	up.PageTotal = CalculatePageCount(total, size)
	for _, u := range users {
		summary := UserSummary{
			Email:    u.EmailAddress,
			Id:       int(u.ID),
			Username: u.Username,
			Active:   !u.Inactive,
		}
		up.Users = append(up.Users, summary)
	}
	return ctx.JSON(http.StatusOK, up)
}

func (c *CerberusAPI) CreateUser(ctx echo.Context) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) InitiatePasswordRecovery(ctx echo.Context) error {
	urp := &UserRecoverPassword{}
	err := ctx.Bind(urp)
	if err != nil {
		return err
	}
	user, err := c.UserStoreService.FindUserByUsername(ctx.Request().Context(), urp.Username)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	otp, err := c.UserStoreService.GenerateUserOTP(ctx.Request().Context(), user.ID, 6)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	println(otp) //TODO integrate communication
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) ResetUserPassword(ctx echo.Context) error {
	urp := &UserResetPassword{}
	err := ctx.Bind(urp)
	if err != nil {
		return err
	}
	user, err := c.UserStoreService.FindUserByUsername(ctx.Request().Context(), urp.Username)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	err = c.UserStoreService.ValidateOTP(ctx.Request().Context(), user.ID, urp.Otp)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	err = c.UserStoreService.SetPassword(ctx.Request().Context(), user.ID, urp.Password)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
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
