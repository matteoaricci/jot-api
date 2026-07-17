package middleware

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/auth"
)

func AuthMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Path()

			if strings.Contains(path, "/public/") {
				return next(c)
			}

			cookie, err := c.Cookie("token")
			if err != nil {
				slog.Warn("auth failed: missing token",
					slog.String("method", c.Request().Method),
					slog.String("path", path),
				)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			token, err := jwt.ParseWithClaims(
				cookie.Value,
				&auth.JwtCustomClaims{},
				func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid signing method")
					}
					return []byte(jwtSecret), nil
				},
			)

			if err != nil {
				slog.Warn("auth failed: invalid token",
					slog.String("method", c.Request().Method),
					slog.String("path", path),
					slog.String("error", err.Error()),
				)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			if !token.Valid {
				slog.Warn("auth failed: token not valid",
					slog.String("method", c.Request().Method),
					slog.String("path", path),
				)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			claims, ok := token.Claims.(*auth.JwtCustomClaims)
			if !ok {
				slog.Warn("auth failed: invalid token claims",
					slog.String("method", c.Request().Method),
					slog.String("path", path),
				)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			userID, err := strconv.ParseUint(claims.UserID, 10, 64)
			if err != nil {
				slog.Warn("auth failed: invalid user id",
					slog.String("method", c.Request().Method),
					slog.String("path", path),
					slog.String("error", err.Error()),
				)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			c.Set("userId", userID)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}
