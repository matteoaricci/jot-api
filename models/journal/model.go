package models

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	CreateOrPutJournalVM struct {
		Title       string `json:"title" validate:"required,max=256"`
		Description string `json:"description" validate:"required,max=256"`
	}

	CreateOrPutJournalValidator struct {
		Validator *validator.Validate
	}
)

func (cjv *CreateOrPutJournalValidator) Validate(i interface{}) error {
	if err := cjv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

type JournalVM struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ID          string `json:"id"`
}
