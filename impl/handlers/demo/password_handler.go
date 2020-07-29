package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandlePasswordPage(c echo.Context) error {
	hpd := NewHomePageData("Resource Owner Password Demo")
	hpd.SetStateCookie(c)
	return c.Render(http.StatusOK, "demo_password.html", hpd)
}

func HandlePasswordPost(c echo.Context) error {
	hpd := NewHomePageData("Resource Owner Password Demo")

	username := c.FormValue("username")
	password := c.FormValue("password")

	token, err := conf.PasswordCredentialsToken(context.Background(), username, password)
	if err != nil {
		hpd.Type = err.Error()
	} else {
		hpd.AccessToken = token.AccessToken
		hpd.RefreshToken = token.RefreshToken
	}

	hpd.SetStateCookie(c)
	return c.Render(http.StatusOK, "demo_home.html", hpd)
}
