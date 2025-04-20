package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/gorm"
	"net/http"
)

func AuthenticateUser(email string, password string) *echo.HTTPError {
	_, err := repo.FindUser(email, password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func SignUpUser(firstName string, lastName string, email string, password string) *echo.HTTPError {
	return nil
}
