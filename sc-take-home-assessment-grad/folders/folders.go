package folders

import (
	"github.com/gofrs/uuid"
)

// What the code does
// The GetAllFolders function is designed to fetch all folders associated with an organization ID
// provided in the FetchFolderRequest and return them in a FetchFolderResponse.
// Function initialises variables
// err,
// f1: Folder struct
// fs: Slice of pointers to Folder structs
// f: a slice of Folder structs
// Fetches folders using the provided OrgID from the request, which returns a slice of pointers
// to Folder structs r and an error
// Function converts the slice of folder pointers r into a slcie of Folder structs
// Functions converts the slice of Folder structs f back into a slice of pointers fp
// Constructs a FetchFolderResponse containing the slice of folder pointers fp and returns it
// The FetchAllFoldersByOrgID function:
// Retrieves a sample set of folders using GetSampleData.
// Filters the folders based on the provided organization ID (OrgID).
// Returns the filtered list of folders.

// Suggested Improvements
// The conversion from a slice of pointers to a slice of structs (f) and then back to a slice of pointers (fp) is redundant. We can skip the intermediate conversion to a slice of structs.
// The error returned by FetchAllFoldersByOrgID is currently ignored. It should be handled appropriately.
// The unnecessary variable initializations (f1, fs) can be removed to clean up the code.
// FetchFolderRequest represents the request to fetch folders

func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	// Fetch all folders by OrgID
	r, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	// Directly use the fetched folder pointers to construct the response
	ffr := &FetchFolderResponse{Folders: r}
	return ffr, nil
}

func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
