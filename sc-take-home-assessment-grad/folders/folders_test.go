package folders_test

import (
	"testing"

	"github.com/amandamcg1/SafetyCultureAssessment/sc-take-home-assessment-grad/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	OrgID1 := folders.GetSampleData()[20].OrgId
	OrgID2, err := uuid.NewV4()

	if err != nil {
		t.Fatalf("failed to generate new IIOD: %v", err)
	}

	t.Run("Successful retrieval of folders for a valid OrgID", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: OrgID1}
		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err, "Expected no error when fetching folders")
		assert.NotNil(t, resp, "Expected a non-nil response")

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

	t.Run("Handle nil request input", func(t *testing.T) {
		req := &folders.FetchFolderRequest{OrgID: OrgID2}
		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err, "Expected no error when no folders are found")
		assert.NotNil(t, resp, "expected a non-nil response")
		assert.Equal(t, 0, len(resp.Folders))
	})

	t.Run("Handle nil request input", func(t *testing.T) {
		resp, err := folders.GetAllFolders(nil)

		assert.Error(t, err, "Expected an error when request is nil")
		assert.Nil(t, resp, "Expected a nil response when request is nil")
	})

	t.Run("Empty OrgID", func(t *testing.T) {
		emptyOrgID := uuid.UUID{} // Represents an empty OrgID
		req := &folders.FetchFolderRequest{OrgID: emptyOrgID}

		resp, err := folders.GetAllFolders(req)

		assert.NoError(t, err, "Expected no error when OrgID is empty")
		assert.Equal(t, 0, len(resp.Folders), "Expected zero folders for an empty OrgID")
	})

}
