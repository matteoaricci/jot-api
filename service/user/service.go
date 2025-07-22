package user

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"github.com/matteoaricci/jot-api/service/journal"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func AuthenticateUser(email string, password string) (*string, *echo.HTTPError) {
	u, err := repo.FindUser(email, password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound)
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	claims := &jwtCustomClaims{
		Name: fmt.Sprintf("%s %s", u.FirstName, u.LastName),
		Role: u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &t, nil
}

// SignUpUser creates a new user with the specified role (defaulting to 'user').
func SignUpUser(firstName string, lastName string, email string, password string, role string) *echo.HTTPError {
	if role == "" {
		role = "user"
	}
	err := repo.CreateUser(firstName, lastName, email, password, role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

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
