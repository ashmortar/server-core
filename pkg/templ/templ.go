package templ

import (
	"core/pkg/config"
	"core/pkg/htmx"
	"core/pkg/layouts"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, content templ.Component) error {
	config := c.Get(config.ContextKey).(config.Config)

	if !htmx.IsHtmxRequest(c) {
		content = layouts.HtmlDoc(config.AppTitle, content)
	}
	return content.Render(
		c.Request().Context(),
		c.Response().Writer,
	)
}
