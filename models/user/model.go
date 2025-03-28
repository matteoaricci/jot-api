package models

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthenticateUserVM struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type AuthenticateUserValidator struct {
	Validator *validator.Validate
}

func (auv *AuthenticateUserValidator) Validate(i interface{}) error {
	if err := auv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
