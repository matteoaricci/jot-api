package entry

import (
	"strconv"

	entryModels "github.com/matteoaricci/jot-api/models/entry"
	"github.com/matteoaricci/jot-api/repo"
)

// MapRepoToVM maps a repo.Entry to its API view model
func MapRepoToVM(e repo.Entry) entryModels.EntryVM {
	return entryModels.EntryVM{
		ID:        strconv.FormatUint(e.ID, 10),
		Content:   e.Content,
		JournalID: e.JournalID,
	}
}

// MapRepoSliceToVMSlice maps a slice of repo.Entry to a slice of EntryVM
func MapRepoSliceToVMSlice(es []repo.Entry) []entryModels.EntryVM {
	vms := make([]entryModels.EntryVM, len(es))
	for i, e := range es {
		vms[i] = MapRepoToVM(e)
	}
	return vms
}
