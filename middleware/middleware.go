package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"time"
)

func Create() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CSRF())

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Uh Oh! You Timed Out Bud!",
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

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			formatRequestLog(c, v)
			return nil
		},
		LogLatency:      true,
		LogMethod:       true,
		LogRoutePath:    true,
		LogRequestID:    true,
		LogStatus:       true,
		LogResponseSize: true,
	}))

	return e
}
