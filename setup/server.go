package setup

import (
	"github.com/identityOrg/cerberus/impl/handlers"
	"github.com/identityOrg/cerberus/impl/handlers/demo"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/identityOrg/oidcsdk"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewEchoServer(serverConfig *config.ServerConfig, templates *AppTemplates,
	manager oidcsdk.IManager, handler *handlers.LoginHandler) *echo.Echo {
	e := echo.New()
	e.Debug = serverConfig.Debug

	p := prometheus.NewPrometheus("cerberus", nil)
	p.Use(e)
	e.Use(middleware.Gzip(), middleware.Secure())
	e.Use(middleware.Logger(), middleware.RequestID(), middleware.Recover())
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	//e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	//	Root:   "static",
	//	Browse: false,
	//}))
	e.Use(StaticFromBox)

	e.Renderer = templates
	manager.SetLoginPageHandler(RenderLoginPage)
	manager.SetConsentPageHandler(RenderConsentPage)
	configureProtocolRouted(e, manager)
	handler.Use(e)

	if serverConfig.Demo {
		demo.Setup(e)
	}

	return e
}

func configureProtocolRouted(e *echo.Echo, manager oidcsdk.IManager) {
	e.GET("/keys", echo.WrapHandler(http.HandlerFunc(manager.ProcessKeysEP)))
	e.GET(oidcsdk.UrlOidcDiscovery, echo.WrapHandler(http.HandlerFunc(manager.ProcessDiscoveryEP)))
	oauth2 := e.Group("/oauth2")
	oauth2.GET("/authorize", echo.WrapHandler(http.HandlerFunc(manager.ProcessAuthorizationEP)))
	oauth2.POST("/token", echo.WrapHandler(http.HandlerFunc(manager.ProcessTokenEP)))
	oauth2.POST("/introspection", echo.WrapHandler(http.HandlerFunc(manager.ProcessIntrospectionEP)))
	oauth2.POST("/revocation", echo.WrapHandler(http.HandlerFunc(manager.ProcessRevocationEP)))
	oauth2.GET("/me", echo.WrapHandler(http.HandlerFunc(manager.ProcessUserInfoEP)))
}
