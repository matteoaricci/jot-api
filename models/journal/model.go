package models

import (
	"database/sql/driver"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

var validate *validator.Validate = validator.New()

type (
	CreateOrPutJournalVM struct {
		Title       string      `json:"title" validate:"required"`
		Description string      `json:"description" validate:"required"`
		Completed   IsCompleted `json:"completed" validate:"omitempty,oneof=true false unknown"`
	}
)

func Validate(i interface{}) error {
	if err := validate.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

type JournalVM struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ID          string      `json:"id"`
	Completed   IsCompleted `json:"completed" validate:"omitempty"`
}

type JournalQueryParams struct {
	Completed IsCompleted `query:"completed" validate:"omitempty,oneof=true false unknown"`
}

type IsCompleted string

const (
	True    IsCompleted = "true"
	False   IsCompleted = "false"
	Unknown IsCompleted = "unknown"
)

func (self *IsCompleted) Scan(value interface{}) error {
	*self = IsCompleted(value.(string))
	return nil
}

func (self IsCompleted) Value() (driver.Value, error) {
	return string(self), nil
}
