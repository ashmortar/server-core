package logger

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type (
	Fmt uint8

	MiddleWareArgs struct {
		Level log.Lvl
		Fmt   Fmt
	}
)

const (
	FMT_JSON Fmt = iota
	FMT_TEXT
)

var time = color.New(color.FgCyan).Sprint("${time_rfc3339}")
var uri = color.New(color.FgMagenta).Sprint("${method} - ${path} ${query}")
var end = "returned: ${status} in ${latency_human}"
var faint = color.New(color.Faint).SprintFunc()

func getRequestLogFormat(format Fmt) string {
	if format == FMT_JSON {
		return middleware.DefaultLoggerConfig.Format
	}

	text := fmt.Sprintf("%s%s%s %s %s\n", faint("["), faint(time), faint("]"), faint(uri), faint(end))
	return text
}

func getHeaderFormat(format Fmt) string {
	if format == FMT_JSON {
		return `{"time":"${time_rfc3339}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}`
	}
	return fmt.Sprintf("[%s] ${level} %s", time, faint("${short_file}:${line}"))
}

func RegisterLoggerMiddleware(args *MiddleWareArgs, e *echo.Echo) {
	e.Logger.SetLevel(args.Level)
	e.Logger.SetHeader(getHeaderFormat(args.Fmt))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: getRequestLogFormat(args.Fmt),
	}))
}
