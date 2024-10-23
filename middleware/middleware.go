package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"time"
)

func Create() *echo.Echo {
	e := echo.New()

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      nil,
		ErrorMessage: "Timeout",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			log.Print(c.Path())
		},
		Timeout: 30 * time.Second,
	}))

	e.Use(middleware.RequestID())

	e.Use(middleware.Gzip())

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		LogLevel:  log.ERROR,
	}))

	return e
}
