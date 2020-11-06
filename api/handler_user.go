package api

import (
	"github.com/identityOrg/cerberus/core/models"
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
	println("OTP: ", otp) //TODO integrate communication
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
	err := c.UserStoreService.DeleteUser(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (c *CerberusAPI) GetUser(ctx echo.Context, id int) error {
	user, err := c.UserStoreService.GetUser(ctx.Request().Context(), uint(id))
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	apiUser := c.convertToApiUser(user)
	return ctx.JSON(http.StatusOK, apiUser)
}

func (c *CerberusAPI) UpdateUser(ctx echo.Context, id int) error {
	return ctx.String(http.StatusNotImplemented, "\"Not Implemented\"")
}

func (c *CerberusAPI) ChangeUserPassword(ctx echo.Context, id int) error {
	cpr := &ChangePasswordRequest{}
	err := ctx.Bind(cpr)
	if err != nil {
		return err
	}
	err = c.UserStoreService.ValidatePassword(ctx.Request().Context(), uint(id), cpr.OldPassword)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	err = c.UserStoreService.SetPassword(ctx.Request().Context(), uint(id), cpr.NewPassword)
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) UpdateUserStatus(ctx echo.Context, id int) error {
	update := &StatusUpdate{}
	err := ctx.Bind(update)
	if err != nil {
		return err
	}
	if update.Active {
		err = c.UserStoreService.ActivateUser(ctx.Request().Context(), uint(id))
	} else {
		err = c.UserStoreService.DeactivateUser(ctx.Request().Context(), uint(id))
	}
	if err != nil {
		return &ApiError{
			ErrorCode: "error",
			Message:   err.Error(),
		}
	}
	return ctx.NoContent(http.StatusAccepted)
}

func (c *CerberusAPI) FindUser(ctx echo.Context, params FindUserParams) error {
	panic("implement me")
}

func (c *CerberusAPI) convertToApiUser(user *models.UserModel) *User {
	apiUser := &User{}
	apiUser.Username = user.Username
	apiUser.Active = !user.Inactive
	metadata := user.Metadata
	if metadata != nil {
		name := metadata.GetName()
		if name != "" {
			apiUser.Name = &name
		}
		email := metadata.GetEmail()
		if email != "" {
			apiUser.UserSummary.Email = email
			apiUser.UserContact.Email = &email
		}
		//apiUser.Birthdate
	}
	return apiUser
}
