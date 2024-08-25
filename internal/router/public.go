package router

import (
	"core/internal/handlers"

	"github.com/labstack/echo/v4"
)

func registerPublicRoutes(e *echo.Echo) {
	// Register the routes for the public pages
	e.GET("/", handlers.RootHandler)
	e.GET("/about", handlers.AboutHandler)
	e.GET("/contact", handlers.ContactHandler)
	e.GET("/privacy", handlers.PrivacyHandler)
	e.GET("/tou", handlers.TouHandler)
}
