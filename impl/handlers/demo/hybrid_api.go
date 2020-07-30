package demo

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

func HandleHybridAPI(c echo.Context) error {
	hpd := NewHomePageData("Hybrid Code Flow Response")
	code := c.FormValue("code")
	exchange, err := conf.Exchange(context.Background(), code,
		oauth2.SetAuthURLParam("state", hpd.State),
		oauth2.SetAuthURLParam("scope", strings.Join(conf.Scopes, " ")),
	)
	if err != nil {
		return err
	} else {
		return c.JSON(http.StatusOK, exchange)
	}
}
