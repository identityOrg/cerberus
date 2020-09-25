package setup

import (
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	static        = rice.MustFindBox("../static")
	staticHandler = http.FileServer(static.HTTPBox())
)

func StaticFromRice(inner echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		c.Logger().Debugf("serving static file for %s", path)
		file, err := static.Open(path)
		if err != nil {
			return inner(c)
		} else {
			stat, err := file.Stat()
			if err != nil {
				return inner(c)
			}
			if !stat.IsDir() {
				println(stat.ModTime().String())
				staticHandler.ServeHTTP(c.Response(), c.Request())
				return nil
			} else {
				return inner(c)
			}
		}
	}
}
