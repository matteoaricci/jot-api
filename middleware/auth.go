package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/auth"
	"net/http"
	"strings"
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
				c.Logger().Warnf("Auth failed for %s %s: missing token", c.Request().Method, path)
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
				c.Logger().Warnf("Auth failed for %s %s: invalid token - %v", c.Request().Method, path, err)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			if !token.Valid {
				c.Logger().Warnf("Auth failed for %s %s: token not valid", c.Request().Method, path)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			claims, ok := token.Claims.(*auth.JwtCustomClaims)
			if !ok {
				c.Logger().Warnf("Auth failed for %s %s: invalid token claims", c.Request().Method, path)
				return echo.NewHTTPError(http.StatusNotFound)
			}

			c.Set("userId", claims.UserID)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}
