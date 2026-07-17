package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func formatRequestLog(c echo.Context, v middleware.RequestLoggerValues) {
	ctx := c.Request().Context()

	slog.InfoContext(ctx, "Request Status",
		slog.Int("status", v.Status),
		slog.Int64("ms", v.Latency.Milliseconds()),
		slog.Int64("bytes", v.ResponseSize),
	)
}
