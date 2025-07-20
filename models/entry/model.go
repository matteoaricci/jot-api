package models

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate *validator.Validate = validator.New()

// CreateEntryVM represents the payload to create a new entry
type CreateEntryVM struct {
	Content string `json:"content" validate:"required"`
}

// EntryVM represents an entry returned by the API
type EntryVM struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	JournalID string `json:"journalId"`
}

// Validate ensures payload fields adhere to requirements
func ValidateEntry(i interface{}) error {
	if err := validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
