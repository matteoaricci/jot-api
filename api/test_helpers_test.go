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

// GetAuthToken creates a test user and returns an auth token cookie
func GetAuthToken(t *testing.T) *http.Cookie {
	t.Helper()

	// Create a test user
	var b bytes.Buffer
	signUpParams := userModels.SignUpUserVM{
		FirstName: "Test",
		LastName:  "User",
		Email:     "testuser@example.com",
		Password:  "testpassword",
		Role:      "user",
	}
	if err := json.NewEncoder(&b).Encode(signUpParams); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/public/sign-up", &b)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	Server.ServeHTTP(rec, req)

	// Authenticate and get token
	b.Reset()
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
