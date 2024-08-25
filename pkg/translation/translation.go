package translation

import (
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type TranslationMiddlewareArgs struct {
	DefaultLanguage    language.Tag
	SupportedLanguages []language.Tag
	ExcludePaths       []string
}

var ContextKey = "TranslationsPrinter"

func determineLanguage(args *TranslationMiddlewareArgs, c echo.Context) language.Tag {

	form := c.FormValue("lang")
	if form != "" {
		for _, lang := range args.SupportedLanguages {
			if lang.String() == form {
				return lang
			}
		}
	}

	accept := c.Request().Header.Get("Accept-Language")
	tags, weights, err := language.ParseAcceptLanguage(accept)
	if err != nil {
		return args.DefaultLanguage
	}
	for i, tag := range tags {
		weight := weights[i]
		if weight < 0.5 {
			return args.DefaultLanguage
		}
		for _, lang := range args.SupportedLanguages {
			if lang == tag {
				return lang
			}
		}
	}
	return args.DefaultLanguage
}

func RegisterTranslationsMiddleWare(args *TranslationMiddlewareArgs) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			for _, path := range args.ExcludePaths {
				if strings.Contains(c.Request().URL.Path, path) {
					return next(c)
				}
			}

			lang := determineLanguage(args, c)

			p := message.NewPrinter(lang)
			c.Set(ContextKey, p)

			return next(c)
		}
	}
}
