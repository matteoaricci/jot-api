package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/models/auth"
	"github.com/matteoaricci/jot-api/repo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAuthCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true

	return cookie
}

func ReadAuthCookie(c echo.Context) error {
	cookie, err := c.Cookie("auth")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

func AuthenticateUser(ctx context.Context, email string, password string, jwtSecret string) (*string, *echo.HTTPError) {
	slog.InfoContext(ctx, "auth.AuthenticateUser")

	u, err := repo.FindUser(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "auth.AuthenticateUser: user not found")
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		slog.WarnContext(ctx, "auth.AuthenticateUser: password mismatch")
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	claims := &auth.JwtCustomClaims{
		UserID: fmt.Sprintf("%d", u.ID),
		Role:   u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &t, nil
}

func SignUpUser(ctx context.Context, firstName string, lastName string, email string, password string, role string) *echo.HTTPError {
	slog.InfoContext(ctx, "auth.SignUpUser")

	if role == "" {
		role = "user"
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = repo.CreateUser(ctx, firstName, lastName, email, string(hashed), role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "auth.SignUpUser: not found")
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
