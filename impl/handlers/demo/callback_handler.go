package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

func HandleCallback(c echo.Context) error {
	hpd := NewHomePageData("Authorization Code Flow Response")
	hpd.ReadStateCookie(c)
	hpd.AccessCode = c.QueryParam("code")
	hpd.AccessToken = c.QueryParam("token")
	if hpd.AccessCode != "" {
		token, err := conf.Exchange(context.Background(), hpd.AccessCode,
			oauth2.SetAuthURLParam("state", hpd.OldState),
			oauth2.SetAuthURLParam("scope", strings.Join(conf.Scopes, " ")),
		)
		if err != nil {
			hpd.Type = err.Error()
			hpd.SetStateCookie(c)
			return c.Render(http.StatusOK, "demo_home.html", hpd)
		} else {
			hpd.AccessToken = token.AccessToken
			hpd.RefreshToken = token.RefreshToken
		}
	}
	hpd.SetStateCookie(c)
	return c.Render(http.StatusOK, "demo_home.html", hpd)
}
