package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	entryAPI "github.com/matteoaricci/jot-api/api/entries"
	entryModels "github.com/matteoaricci/jot-api/models/entry"
	"github.com/stretchr/testify/assert"
)

// TestEntryEndpoints covers GET and POST for journal entries
func TestEntryEndpoints(t *testing.T) {
	// ensure the entry routes are registered
	_ = entryAPI.AddRoutes

	t.Run("Get entries (none exist)", func(t *testing.T) {
		e := Server
		req := httptest.NewRequest(http.MethodGet, "/api/journals/1/entries", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "[]", rec.Body.String())
	})

	t.Run("Create entry", func(t *testing.T) {
		e := Server

		var b bytes.Buffer
		dummy := entryModels.CreateEntryVM{Content: "test content"}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/journals/1/entries", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "\"1\"\n", rec.Body.String())
	})

	t.Run("Get entries (after create)", func(t *testing.T) {
		e := Server
		req := httptest.NewRequest(http.MethodGet, "/api/journals/1/entries", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)

		var got []entryModels.EntryVM
		if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
			t.Fatal(err)
		}
		assert.Len(t, got, 1)
		assert.Equal(t, "1", got[0].ID)
		assert.Equal(t, "test content", got[0].Content)
		assert.Equal(t, "1", got[0].JournalID)
	})

	t.Run("Create entry missing content", func(t *testing.T) {
		e := Server

		var b bytes.Buffer
		dummy := struct {
			Foo string `json:"foo"`
		}{Foo: "bar"}
		if err := json.NewEncoder(&b).Encode(dummy); err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/api/journals/1/entries", &b)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}
