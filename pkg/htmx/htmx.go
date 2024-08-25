package htmx

import (
	"github.com/labstack/echo/v4"
)

func IsHtmxRequest(h echo.Context) bool {
	return h.Request().Header.Get("HX-Request") == "true"
}

func SetHxReplaceUrl(h echo.Context, url string) {
	h.Response().Header().Set("HX-Replace-Url", url)
}
func HtmlContentTypeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "text/html")
		return next(c)
	}
}
