package user

import (
	"errors"
	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"github.com/matteoaricci/jot-api/service/journal"
	"gorm.io/gorm"
	"net/http"
)

func GetJournals(id string, params models.JournalQueryParams) ([]models.JournalVM, *echo.HTTPError) {
	js, err := repo.GetJournalsByUserID(id, params)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	jVMs := journal.MapRepoSliceToVMSlice(js)

	return jVMs, nil
}

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
