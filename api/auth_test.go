package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	userModels "github.com/matteoaricci/jot-api/models/user"
)

func TestAuthEndpoints(t *testing.T) {
	e := Server

	t.Run("Sign up new user", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.SignUpUserVM{
			FirstName: "Auth",
			LastName:  "Test",
			Email:     "authtest@example.com",
			Password:  "password123",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/public/sign-up", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Authenticate user and receive token cookie", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.AuthenticateUserVM{
			Email:    "authtest@example.com",
			Password: "password123",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/public/authenticate", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)

		cookies := rec.Result().Cookies()
		assert.NotEmpty(t, cookies, "should receive a cookie")

		var tokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "token" {
				tokenCookie = cookie
				break
			}
		}
		assert.NotNil(t, tokenCookie, "should have token cookie")
		assert.NotEmpty(t, tokenCookie.Value, "token cookie should have value")
		assert.True(t, tokenCookie.HttpOnly, "token cookie should be HttpOnly")
		assert.True(t, tokenCookie.Secure, "token cookie should be Secure")
	})

	t.Run("Authenticate with invalid credentials", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.AuthenticateUserVM{
			Email:    "wrong@example.com",
			Password: "wrongpassword",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/public/authenticate", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Authenticate existing user with wrong password", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.AuthenticateUserVM{
			Email:    "authtest@example.com",
			Password: "notthepassword",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/public/authenticate", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Access protected route without token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/journals", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Access protected route with valid token", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.AuthenticateUserVM{
			Email:    "authtest@example.com",
			Password: "password123",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		authReq := httptest.NewRequest(http.MethodPost, "/api/public/authenticate", &b)
		authReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		authRec := httptest.NewRecorder()
		e.ServeHTTP(authRec, authReq)

		cookies := authRec.Result().Cookies()
		var tokenCookie *http.Cookie
		for _, cookie := range cookies {
			if cookie.Name == "token" {
				tokenCookie = cookie
				break
			}
		}
		assert.NotNil(t, tokenCookie)

		req := httptest.NewRequest(http.MethodGet, "/api/journals", nil)
		req.AddCookie(tokenCookie)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Access protected route with invalid token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/journals", nil)
		req.AddCookie(&http.Cookie{
			Name:  "token",
			Value: "invalid.token.here",
		})
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Public routes accessible without token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/public/healthz", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
