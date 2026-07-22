package stats

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/matteoaricci/jot-api/models/stats"
	"github.com/matteoaricci/jot-api/repo"
)

func GetSystemStats(ctx context.Context) (*models.SystemStatsVM, *echo.HTTPError) {
	slog.InfoContext(ctx, "stats.GetSystemStats")

	s, err := repo.GetSystemStats(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &models.SystemStatsVM{
		TotalUsers:        s.TotalUsers,
		TotalJournals:     s.TotalJournals,
		TotalEntries:      s.TotalEntries,
		CompletedJournals: s.CompletedJournals,
	}, nil
}

func GetUserStats(ctx context.Context) ([]models.UserStatsVM, *echo.HTTPError) {
	slog.InfoContext(ctx, "stats.GetUserStats")

	rows, err := repo.GetUserStats(ctx)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	vms := make([]models.UserStatsVM, 0, len(rows))
	for _, r := range rows {
		vms = append(vms, models.UserStatsVM{
			UserID:       r.UserID,
			Email:        r.Email,
			JournalCount: r.JournalCount,
			EntryCount:   r.EntryCount,
		})
	}

	return vms, nil
}
