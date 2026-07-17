package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/gorm"
	"net/http"
)

func DeleteUser(id string) *echo.HTTPError {
	err := repo.DeleteUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
