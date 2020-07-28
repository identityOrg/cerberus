package setup

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewServer() *echo.Echo {
	e := echo.New()

	p := prometheus.NewPrometheus("cerberus", nil)
	p.Use(e)
	e.Use(middleware.Gzip(), middleware.Secure())
	e.Use(middleware.Logger(), middleware.RequestID(), middleware.Recover())
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: false,
	}))

	e.Renderer = NewAppTemplates()

	manager, err := NewManager()
	if err != nil {
		panic(err)
	}
	ConfigureEcho(e, manager)

	return e
}
