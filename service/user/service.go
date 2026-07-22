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

func UpdateRole(ctx context.Context, userID uint64, roleName string) *echo.HTTPError {
	slog.InfoContext(ctx, "user.UpdateRole", slog.Uint64("userId", userID), slog.String("role", roleName))

	role, err := repo.FindRoleByName(ctx, roleName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "user.UpdateRole: unknown role", slog.String("role", roleName))
			return echo.NewHTTPError(http.StatusBadRequest, "unknown role")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := repo.UpdateUserRole(ctx, userID, role.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.WarnContext(ctx, "user.UpdateRole: user not found", slog.Uint64("userId", userID))
			return echo.NewHTTPError(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

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
