package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEndpoints(t *testing.T) {
	t.Run("health check endpoint should return 200", func(t *testing.T) {
		e := echo.New()

		AddRouteHandlers(e)

		req := httptest.NewRequest(http.MethodGet, "/api/healthz", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, `{"status":"OK"}`+"\n", rec.Body.String())
		assert.Equal(t, rec.Header().Get("Content-Type"), "application/json")
	})

	t.Run("unknown route should return 404", func(t *testing.T) {
		e := echo.New()

		AddRouteHandlers(e)

		req := httptest.NewRequest(http.MethodGet, "/route-not-found", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	})
}
