package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	userModels "github.com/matteoaricci/jot-api/models/user"
)

// GetAuthToken authenticates as the seeded test user (usr id 1, who owns the
// seeded journals) and returns an auth token cookie.
func GetAuthToken(t *testing.T) *http.Cookie {
	t.Helper()

	// Authenticate as the user seeded in testdata/integration_test_init.sql
	var b bytes.Buffer
	authParams := userModels.AuthenticateUserVM{
		Email:    "testuser@example.com",
		Password: "testpassword",
	}
	if err := json.NewEncoder(&b).Encode(authParams); err != nil {
		t.Fatal(err)
	}

	authReq := httptest.NewRequest(http.MethodPost, "/api/public/authenticate", &b)
	authReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	authRec := httptest.NewRecorder()
	Server.ServeHTTP(authRec, authReq)

	cookies := authRec.Result().Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "token" {
			return cookie
		}
	}

	t.Fatal("no token cookie found")
	return nil
}
