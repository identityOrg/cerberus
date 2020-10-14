package demo

import (
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	conf1       *oauth2.Config
	clientConf1 *clientcredentials.Config
)

func getAuthCodeConfig() *oauth2.Config {
	if conf1 != nil {
		return conf1
	}
	conf1 = &oauth2.Config{
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
	sdkConfig := config.NewSDKConfig()
	conf1.RedirectURL = sdkConfig.Issuer + "/redirect"
	conf1.Endpoint.AuthURL = sdkConfig.Issuer + "/oauth2/authorize"
	conf1.Endpoint.TokenURL = sdkConfig.Issuer + "/oauth2/token"
	return conf1
}

func getClientCredConfig() *clientcredentials.Config {
	if clientConf1 != nil {
		return clientConf1
	}
	clientConf1 = &clientcredentials.Config{
		ClientID:     "client",
		ClientSecret: "client",
		TokenURL:     "http://localhost:8080/oauth2/token",
		Scopes:       []string{"openid", "offline", "offline_access"},
		AuthStyle:    oauth2.AuthStyleInHeader,
	}
	sdkConfig := config.NewSDKConfig()
	clientConf1.TokenURL = sdkConfig.Issuer + "/oauth2/token"
	return clientConf1
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
