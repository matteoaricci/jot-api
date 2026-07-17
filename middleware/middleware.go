package middleware

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddMiddleware(e *echo.Echo, jwtSecret string) {
	//e.Use(middleware.CSRF())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "Uh Oh! You Timed Out Bud!",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			slog.Warn("request timed out", slog.String("uri", c.Request().RequestURI))
		},
		Timeout: 0 * time.Second,
	}))

	e.Use(middleware.RequestID())

	e.Use(middleware.Gzip())

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			slog.Error("recovered from panic",
				slog.String("error", err.Error()),
				slog.String("stack", string(stack)),
			)
			return err
		},
	}))

	e.Use(AuthMiddleware(jwtSecret))

	e.Use(LogContextMiddleware())

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
}
