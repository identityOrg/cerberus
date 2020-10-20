package setup

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/identityOrg/cerberus/api"
	"github.com/identityOrg/cerberus/impl/handlers"
	"github.com/identityOrg/cerberus/impl/handlers/demo"
	"github.com/identityOrg/cerberus/setup/config"
	"github.com/identityOrg/cerberus/setup/middle"
	"github.com/identityOrg/oidcsdk"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewEchoServer(serverConfig *config.ServerConfig, templates *AppTemplates, config2 *oidcsdk.Config,
	manager oidcsdk.IManager, handler *handlers.LoginHandler, si api.ServerInterface) *echo.Echo {
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
	e.Use(StaticFromRice)

	e.Renderer = templates
	manager.SetLoginPageHandler(RenderLoginPage)
	manager.SetConsentPageHandler(RenderConsentPage)
	configureProtocolRouted(e, manager)
	handler.Use(e)

	if serverConfig.Demo {
		demo.Setup(e)
	}

	if serverConfig.Api {
		api.RegisterHandlers(e, si)
		e.GET("/v1/spec", func(context echo.Context) error {
			swagger, err := api.GetSwagger()
			if err != nil {
				return err
			}
			s := &openapi3.Server{
				URL:         config2.Issuer,
				Description: "Primary Server",
			}
			swagger.Servers = append(swagger.Servers, s)
			return context.JSON(http.StatusOK, swagger)
		})
	}

	return e
}

func configureProtocolRouted(e *echo.Echo, manager oidcsdk.IManager) {
	e.GET(oidcsdk.UrlOidcDiscovery, echo.WrapHandler(http.HandlerFunc(manager.ProcessDiscoveryEP)))
	oauth2 := e.Group("/oauth2", middle.NoCache())
	oauth2.GET("/authorize", echo.WrapHandler(http.HandlerFunc(manager.ProcessAuthorizationEP)))
	oauth2.POST("/token", echo.WrapHandler(http.HandlerFunc(manager.ProcessTokenEP)))
	oauth2.POST("/introspection", echo.WrapHandler(http.HandlerFunc(manager.ProcessIntrospectionEP)))
	oauth2.POST("/revocation", echo.WrapHandler(http.HandlerFunc(manager.ProcessRevocationEP)))
	oauth2.GET("/me", echo.WrapHandler(http.HandlerFunc(manager.ProcessUserInfoEP)))
	oauth2.GET("/keys", echo.WrapHandler(http.HandlerFunc(manager.ProcessKeysEP)))
}
