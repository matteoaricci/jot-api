package user

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matteoaricci/jot-api/repo"
	"gorm.io/gorm"
)

func DeleteUser(ctx context.Context, id string) *echo.HTTPError {
	slog.InfoContext(ctx, "user.DeleteUser", slog.String("id", id))

	err := repo.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "user.DeleteUser: not found", slog.String("id", id))
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}
