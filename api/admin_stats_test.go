package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	statsModels "github.com/matteoaricci/jot-api/models/stats"
	"github.com/stretchr/testify/assert"
)

func TestAdminStats(t *testing.T) {
	e := Server

	get := func(t *testing.T, path string, cookie *http.Cookie) *httptest.ResponseRecorder {
		t.Helper()
		req := httptest.NewRequest(http.MethodGet, path, nil)
		if cookie != nil {
			req.AddCookie(cookie)
		}
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)
		return rec
	}

	t.Run("system stats", func(t *testing.T) {
		t.Run("no token is rejected", func(t *testing.T) {
			assert.Equal(t, http.StatusNotFound, get(t, "/api/admin/stats", nil).Code)
		})

		t.Run("non-admin is forbidden", func(t *testing.T) {
			assert.Equal(t, http.StatusForbidden, get(t, "/api/admin/stats", GetAuthToken(t)).Code)
		})

		t.Run("admin gets system counts", func(t *testing.T) {
			rec := get(t, "/api/admin/stats", GetAdminAuthToken(t))
			assert.Equal(t, http.StatusOK, rec.Code)

			var s statsModels.SystemStatsVM
			if err := json.Unmarshal(rec.Body.Bytes(), &s); err != nil {
				t.Fatal(err)
			}
			// Seed: 2 users, 3 journals (1 completed), 0 entries.
			assert.Equal(t, int64(2), s.TotalUsers)
			assert.Equal(t, int64(3), s.TotalJournals)
			assert.Equal(t, int64(1), s.CompletedJournals)
			assert.Equal(t, int64(0), s.TotalEntries)
		})
	})

	t.Run("user stats", func(t *testing.T) {
		t.Run("non-admin is forbidden", func(t *testing.T) {
			assert.Equal(t, http.StatusForbidden, get(t, "/api/admin/users/stats", GetAuthToken(t)).Code)
		})

		t.Run("admin gets per-user counts", func(t *testing.T) {
			rec := get(t, "/api/admin/users/stats", GetAdminAuthToken(t))
			assert.Equal(t, http.StatusOK, rec.Code)

			var rows []statsModels.UserStatsVM
			if err := json.Unmarshal(rec.Body.Bytes(), &rows); err != nil {
				t.Fatal(err)
			}
			// One row per seeded user.
			assert.Len(t, rows, 2)

			byID := map[uint64]statsModels.UserStatsVM{}
			for _, r := range rows {
				byID[r.UserID] = r
			}
			// User 1 owns all 3 seeded journals; user 2 (admin) owns none.
			assert.Equal(t, int64(3), byID[1].JournalCount)
			assert.Equal(t, int64(0), byID[2].JournalCount)
		})
	})
}
