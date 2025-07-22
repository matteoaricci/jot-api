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

func TestUserEndpoints(t *testing.T) {
	e := Server

	t.Run("Sign up user", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.SignUpUserVM{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@example.com",
			Password:  "secret",
			Role:      "admin",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/sign-up", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("Authenticate user", func(t *testing.T) {
		var b bytes.Buffer
		dummy := userModels.AuthenticateUserVM{
			Email:    "jane@example.com",
			Password: "secret",
		}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/api/authenticate", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "token")

		// invalid credentials
		dummy = userModels.AuthenticateUserVM{Email: "nope@example.com", Password: "bad"}
		b.Reset()
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}
		req = httptest.NewRequest(http.MethodPost, "/api/authenticate", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("Get User Journals", func(t *testing.T) {
		t.Run("No Params", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/users/1/journals", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t, `[]`, rec.Body.String())
		})

		t.Run("Invalid Params", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/users/1/journals?size=foo", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		})
	})

	t.Run("Delete user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/users/1", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
