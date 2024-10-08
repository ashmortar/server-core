package handlers

import (
	"core/pkg/htmx"
	"core/pkg/templ"
	"core/pkg/translation"

	"core/pkg/pages"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/message"
)

func RootHandler(c echo.Context) error {
	if htmx.IsHtmxRequest(c) {
		htmx.SetHxPushUrl(c, "/")
	}
	return templ.Render(
		c, pages.Root(c.Get(translation.ContextKey).(*message.Printer)),
	)
}

func AboutHandler(c echo.Context) error {
	if htmx.IsHtmxRequest(c) {
		htmx.SetHxPushUrl(c, "/about")
	}
	return templ.Render(
		c, pages.About(c.Get(translation.ContextKey).(*message.Printer)),
	)
}

func ContactHandler(c echo.Context) error {
	if htmx.IsHtmxRequest(c) {
		htmx.SetHxPushUrl(c, "/contact")
	}
	return templ.Render(
		c, pages.Contact(c.Get(translation.ContextKey).(*message.Printer)),
	)
}

func PrivacyHandler(c echo.Context) error {
	if htmx.IsHtmxRequest(c) {
		htmx.SetHxPushUrl(c, "/privacy")
	}
	return templ.Render(c, pages.Privacy())
}

func TouHandler(c echo.Context) error {
	if htmx.IsHtmxRequest(c) {
		htmx.SetHxPushUrl(c, "/tou")
	}
	return templ.Render(c, pages.Tou())
}
