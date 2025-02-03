package journal

import (
	models "github.com/matteoaricci/jot-api/models/journal"
	"github.com/matteoaricci/jot-api/repo"
	"strconv"
)

func MapRepoToVM(j repo.Journal) models.JournalVM {
	return models.JournalVM{
		Title:       j.Title,
		Description: j.Description,
		ID:          strconv.FormatUint(j.ID, 10),
		Completed:   j.Completed,
	}
}

func MapRepoSliceToVMSlice(js []repo.Journal) []models.JournalVM {
	jVMs := make([]models.JournalVM, 0)
	for _, j := range js {
		jVMs = append(jVMs, MapRepoToVM(j))
	}

	return jVMs
}

func RepoToPageOfVMs(js []repo.Journal, params models.JournalQueryParams) models.PageOfJournalVMs {
	total := len(js)

	pageOfVMs := models.PageOfJournalVMs{
		TotalRecords: total,
		Journals:     MapRepoSliceToVMSlice(js),
		Page:         params.Page,
		Size:         params.Size,
	}

	return pageOfVMs
}
