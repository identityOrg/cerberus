package middle

import "github.com/labstack/echo/v4"

func NoCache() echo.MiddlewareFunc {
	return func(inner echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cache-Control", "no-store, must-revalidate")
			return inner(c)
		}
	}
}
