package router

import (
	"github.com/labstack/echo/v4"
)

func RegisterServerRoutes(e *echo.Echo) {
	registerPublicRoutes(e)
}
