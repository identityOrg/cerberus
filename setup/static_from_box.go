package setup

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"net/http"
)

func StaticFromBox(h echo.HandlerFunc) echo.HandlerFunc {
	box := packr.New("static", "../static")
	fileServer := http.FileServer(box)
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if box.Has(path) {
			c.Logger().Debugf("serving static file for %s", path)
			fileServer.ServeHTTP(c.Response(), c.Request())
			return nil
		} else {
			return h(c)
		}
	}
}
