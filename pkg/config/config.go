package config

import (
	"core/pkg/logger"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var ContextKey = "Config"

var validate = validator.New(validator.WithRequiredStructEnabled())

var yellowsp = color.New(color.FgYellow).SprintFunc()

var red = color.New(color.FgRed)
var redSp = red.SprintFunc()

type Config struct {
	DatabaseURL string     `validate:"required"`
	Port        int        `validate:"required"`
	LogLevel    log.Lvl    `validate:"oneof=1 2 3 4 5"`
	LogFmt      logger.Fmt `validate:"oneof=0 1"`
	AppTitle    string     `validate:"required"`
}

var config Config

func check() (Config, error) {
	err := validate.Struct(config)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errMsg := "Config validation failed:\n"
		for _, e := range validationErrors {
			errMsg += fmt.Sprintf("%v\n", redSp(e.Error()))
			val := e.Value()
			param := e.Param()
			if param != "" {
				errMsg += red.Sprintf("\treceived: %v\n\texpected: %v\n", yellowsp(val), yellowsp(param))
			}
		}
		return config, errors.New(errMsg)
	}
	return config, nil
}

func Parse() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("LOG_LEVEL", log.DEBUG)
	viper.SetDefault("LOG_FMT", logger.FMT_JSON)
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	viper.AutomaticEnv()
	config = Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		Port:        viper.GetInt("PORT"),
		LogLevel:    log.Lvl(viper.GetInt("LOG_LEVEL")),
		LogFmt:      logger.Fmt(viper.GetInt("LOG_FMT")),
		AppTitle:    viper.GetString("APP_TITLE"),
	}

	config, err = check()
	if err != nil {
		return config, err
	}

	return config, nil
}

func RegisterConfigMiddleware(conf Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Logger().SetLevel(conf.LogLevel)
			c.Set(ContextKey, conf)
			return next(c)
		}
	}
}
