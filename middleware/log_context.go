package middleware

import (
	"log/slog"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/logger"
)

// LogContextMiddleware stashes request-scoped attributes into the request
// context. Runs after Auth so userId is available.
func LogContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()

			userID := "unknown"
			if id, ok := c.Get("userId").(uint64); ok {
				userID = strconv.FormatUint(id, 10)
			}

			ctx := logger.WithAttrs(req.Context(),
				slog.Group("http",
					slog.String("method", req.Method),
					slog.String("route", c.Path()),
					slog.String("remoteAddress", req.RemoteAddr),
					slog.String("X-Forwarded-For", req.Header.Get(echo.HeaderXForwardedFor)),
				),
				slog.String("correlationId", c.Response().Header().Get(echo.HeaderXRequestID)),
				slog.String("userId", userID),
			)

			c.SetRequest(req.WithContext(ctx))

			slog.InfoContext(ctx, "Incoming Request", slog.Int64("bytes", req.ContentLength))

			return next(c)
		}
	}
}
