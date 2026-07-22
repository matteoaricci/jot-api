package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserRole(t *testing.T) {
	e := Server

	putRole := func(t *testing.T, cookie *http.Cookie, targetID string, body string) *httptest.ResponseRecorder {
		t.Helper()
		req := httptest.NewRequest(http.MethodPut, "/api/admin/users/"+targetID+"/role", bytes.NewBufferString(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		if cookie != nil {
			req.AddCookie(cookie)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		return rec
	}

	t.Run("no token is rejected", func(t *testing.T) {
		rec := putRole(t, nil, "1", `{"role":"admin"}`)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("non-admin is forbidden", func(t *testing.T) {
		rec := putRole(t, GetAuthToken(t), "1", `{"role":"admin"}`)
		assert.Equal(t, http.StatusForbidden, rec.Code)
	})

	t.Run("admin can update a user's role", func(t *testing.T) {
		rec := putRole(t, GetAdminAuthToken(t), "1", `{"role":"admin"}`)
		assert.Equal(t, http.StatusNoContent, rec.Code)

		// Re-authenticating as user 1 now yields an admin role claim.
		var b bytes.Buffer
		_ = json.NewEncoder(&b).Encode(map[string]string{
			"email":    "testuser@example.com",
			"password": "testpassword",
		})
		authReq := httptest.NewRequest(http.MethodPost, "/api/public/authenticate", &b)
		authReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		authRec := httptest.NewRecorder()
		e.ServeHTTP(authRec, authReq)
		assert.Equal(t, http.StatusOK, authRec.Code)

		// restore user 1 back to the user role so other tests are unaffected
		restore := putRole(t, GetAdminAuthToken(t), "1", `{"role":"user"}`)
		assert.Equal(t, http.StatusNoContent, restore.Code)
	})

	t.Run("unknown role is a bad request", func(t *testing.T) {
		rec := putRole(t, GetAdminAuthToken(t), "1", `{"role":"wizard"}`)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("invalid id is a bad request", func(t *testing.T) {
		rec := putRole(t, GetAdminAuthToken(t), "abc", `{"role":"user"}`)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("missing role is a bad request", func(t *testing.T) {
		rec := putRole(t, GetAdminAuthToken(t), "1", `{}`)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
