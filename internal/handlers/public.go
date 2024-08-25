package handlers

import (
	"core/pkg/templ"
	"core/pkg/translation"

	"core/pkg/pages"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/message"
)

func RootHandler(c echo.Context) error {
	return templ.Render(
		c, pages.Root(c.Get(translation.ContextKey).(*message.Printer)),
	)
}

func AboutHandler(c echo.Context) error {
	return templ.Render(c, pages.About())
}

func ContactHandler(c echo.Context) error {
	return templ.Render(c, pages.Contact())
}

func PrivacyHandler(c echo.Context) error {
	return templ.Render(c, pages.Privacy())
}

func TouHandler(c echo.Context) error {
	return templ.Render(c, pages.Tou())
}
