package server

import (
	"context"
	"core/internal/middleware"
	"core/internal/router"
	"core/pkg/config"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	_ "core/internal/translations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Init() {
	// validate env
	conf, err := config.Parse()
	if err != nil {
		fmt.Printf("%e", err)
	}
	// Initialize Echo
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = conf.LogLevel == log.DEBUG
	// Middleware
	middleware.Init(e, conf)
	// Register routes
	router.RegisterServerRoutes(e)

	// Start server
	go func() {

		if err := e.Start(":" + strconv.Itoa(conf.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
