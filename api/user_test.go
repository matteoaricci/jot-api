package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserEndpoints(t *testing.T) {
	e := Server

	t.Run("Get User Journals", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users/:id/journals", nil)
		res := httptest.NewRecorder()

		e.ServeHTTP(req, res)
	})

}
