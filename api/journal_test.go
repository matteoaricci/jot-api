package api

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/api/journals"
	"github.com/matteoaricci/jot-api/models/journal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJournalEndpoints(t *testing.T) {
	t.Run("Get all journals", func(t *testing.T) {
		e := Server

		t.Run("No Params", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/journals", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{
			  "total_records" : 3,
			  "journals" : [ {
				"title" : "Psychopomp",
				"description" : "Japanese Breakfast's first album",
				"id" : "1",
				"completed" : "true"
			  }, {
				"title" : "Soft Sounds from Another Planet",
				"description" : "Absolute banger followup",
				"id" : "2",
				"completed" : "false"
			  }, {
				"title" : "Jubilee",
				"description" : "Here Michelle Zauner asks: what if joy was as complex as grief",
				"id" : "3",
				"completed" : "unknown"
			  } ],
			  "page" : 1,
			  "size" : 10
			}`,
				rec.Body.String())
		})

		t.Run("With Pagination Params", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/journals?size=1&page=2", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{
			  "total_records" : 1,
			  "journals" : [{
				"title" : "Soft Sounds from Another Planet",
				"description" : "Absolute banger followup",
				"id" : "2",
				"completed" : "false"
			  }],
			  "page" : 2,
			  "size" : 1
			}`,
				rec.Body.String())
		})

		t.Run("With Completed Params", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/journals?completed=unknown", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{
			  "total_records" : 1,
			  "journals" : [{
				"title" : "Jubilee",
				"description" : "Here Michelle Zauner asks: what if joy was as complex as grief",
				"id" : "3",
				"completed" : "unknown"
			  } ],
			  "page" : 1,
			  "size" : 10
			}`,
				rec.Body.String())
		})
	})
	t.Run("Get journal by id", func(t *testing.T) {
		e := Server

		t.Run("success", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/journals/1", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{"title":"Psychopomp","description":"Japanese Breakfast's first album","id":"1","completed":"true"}`,
				rec.Body.String())
		})

		t.Run("not found", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/journals/4", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusNotFound, rec.Code)
		})
	})

	t.Run("Create journal", func(t *testing.T) {
		e := Server

		t.Run("success", func(t *testing.T) {
			var b bytes.Buffer

			dummyData := models.CreateOrPutJournalVM{
				Title:       "dummy title",
				Description: "dummy desc",
				Completed:   "true",
			}

			err := json.NewEncoder(&b).Encode(dummyData)

			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/journals", &b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "\"4\"\n", rec.Body.String())
		})

		t.Run("missing data", func(t *testing.T) {
			var b bytes.Buffer

			dummyData := struct {
				Title string `json:"title"`
			}{
				Title: "dummy title",
			}

			err := json.NewEncoder(&b).Encode(dummyData)

			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/journals", &b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{"message":"Key: 'CreateOrPutJournalVM.Description' Error:Field validation for 'Description' failed on the 'required' tag"}`, rec.Body.String())
		})
	})

	t.Run("Update journal", func(t *testing.T) {
		e := Server

		t.Run("success", func(t *testing.T) {
			var b bytes.Buffer

			dummyData := models.CreateOrPutJournalVM{
				Title:       "dummy title",
				Description: "dummy desc",
				Completed:   "false",
			}

			err := json.NewEncoder(&b).Encode(dummyData)

			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(http.MethodPut, "/api/journals/2", &b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{"title":"dummy title","description":"dummy desc", "id":  "2","completed":"false"}`,
				rec.Body.String())
		})

		t.Run("not found", func(t *testing.T) {
			var b bytes.Buffer

			dummyData := models.CreateOrPutJournalVM{
				Title:       "dummy title",
				Description: "dummy desc",
			}

			err := json.NewEncoder(&b).Encode(dummyData)

			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(http.MethodPut, "/api/journals/20", &b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusNotFound, rec.Code)
		})

		t.Run("missing data", func(t *testing.T) {
			var b bytes.Buffer

			dummyData := struct {
				Title string `json:"title"`
			}{
				Title: "dummy title",
			}

			err := json.NewEncoder(&b).Encode(dummyData)

			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(http.MethodPut, "/api/journals/2", &b)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.JSONEq(t,
				// language=JSON
				`{"message":"Key: 'CreateOrPutJournalVM.Description' Error:Field validation for 'Description' failed on the 'required' tag"}`, rec.Body.String())

		})
	})

	t.Run("Delete journal", func(t *testing.T) {
		e := Server
		journals.AddRoutes(e)
		t.Run("success", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/api/journals/1", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusNoContent, rec.Code)

			req = httptest.NewRequest(http.MethodGet, "/api/journals/1", nil)
			rec = httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusNotFound, rec.Code)
		})

		t.Run("not found", func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/api/journals/20", nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, http.StatusNotFound, rec.Code)
		})
	})
}
