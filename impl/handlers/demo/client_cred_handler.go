package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleClientCredential(c echo.Context) error {
	hpd := NewHomePageData("Client Credentials Example")

	token, err := getClientCredConfig().Token(context.Background())
	if err != nil {
		hpd.Type = err.Error()
	} else {
		hpd.AccessToken = token.AccessToken
		hpd.RefreshToken = token.RefreshToken
	}

	return c.Render(http.StatusOK, "demo_home.html", hpd)
}
