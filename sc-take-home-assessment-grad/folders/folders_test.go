package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	// "github.com/georgechieng-sc/interns-2022/folders"
)

func Test_GetAllFolders(t *testing.T) {
	OrgID1 := folders.GetSampleData()[20].OrgId
	OrgID2, err := uuid.NewV4()
	t.Run("Successful retrieval of folders for a valid OrgID", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: OrgID1}
		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err, "Expected no error when fetching folders")

		expectedCount := 1
		assert.Equal(t, expectedCount, len(resp.Folders), "Expected the correct number of folders")

		for _, folder := range resp.Folders {
			assert.Equal(t, OrgID1, folder.OrgId, "Expected OrgID to match")
		}

	})

	t.Run("No folders found for non-existing OrgID", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: OrgID2}
		resp, resperr := folders.GetAllFolders(req)

		assert.NoError(t, err, "No Error createing new Id")

		assert.NoError(t, resperr, "Expected no error when no folders are found")
		assert.Equal(t, 0, len(resp.Folders), "Expected zero folders for a non-existing OrgID")
	})
}
