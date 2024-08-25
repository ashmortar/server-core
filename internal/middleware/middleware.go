package middleware

import (
	"core/internal/assets"
	"core/internal/translations"
	"core/pkg/config"
	"core/pkg/logger"
	"core/pkg/translation"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/text/language"
)

func Init(e *echo.Echo, conf config.Config) {
	e.Use(config.RegisterConfigMiddleware(conf))
	e.Use(middleware.Gzip())
	e.Use(middleware.RequestID())
	logger.RegisterLoggerMiddleware(&logger.MiddleWareArgs{
		Level: conf.LogLevel,
		Fmt:   conf.LogFmt,
	}, e)
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(assets.Files),
		Root:       "",
	}))
	e.Use(translation.RegisterTranslationsMiddleWare(&translation.TranslationMiddlewareArgs{
		DefaultLanguage:    language.English,
		SupportedLanguages: translations.Languages,
		ExcludePaths:       []string{"/css", "/js", "/img", "/fonts"},
	}))
}
