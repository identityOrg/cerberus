package demo

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var conf = &oauth2.Config{
	ClientID:     "client",
	ClientSecret: "client",
	Scopes:       []string{"openid", "offline", "offline_access"},
	Endpoint: oauth2.Endpoint{
		AuthURL:   "http://localhost:8080/oauth2/authorize",
		TokenURL:  "http://localhost:8080/oauth2/token",
		AuthStyle: oauth2.AuthStyleInHeader,
	},
	RedirectURL: "http://localhost:8080/redirect",
}

var clientConf = clientcredentials.Config{
	ClientID:     "client",
	ClientSecret: "client",
	TokenURL:     "http://localhost:8080/oauth2/token",
	Scopes:       []string{"openid", "offline", "offline_access"},
	AuthStyle:    oauth2.AuthStyleInHeader,
}

func Setup(e *echo.Echo) {
	e.GET("/", HandleHome)
	e.GET("/redirect", HandleCallback)
	e.GET("/password", HandlePasswordPage)
	e.POST("/password", HandlePasswordPost)
	e.GET("/client_credentials", HandleClientCredential)
	e.POST("/refresh", HandleRefresh)
	e.POST("/hybrid", HandleHybridAPI)
}
