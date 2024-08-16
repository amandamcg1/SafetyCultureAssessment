package folders

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/gofrs/uuid"
)

// This file handles paginated folder retrieval. The `PaginatedGetAllFolders` function fetches a specific slice
// of folders based on the request, including the offset and limit. It determines which folders to return and
// creates a new token if there's more data to fetch. `PaginatedFetchAllFoldersByOrgID` simulates getting folders
// filtered by the orgID. Tokens are just encoded integers that help keep track of where you are in the data.
// `generateNextToken` and `DecodeToken` are used to create and read these tokens. This setup helps break down
// large amounts of data into smaller, manageable chunks.

func PaginatedGetAllFolders(req *PaginatedFetchFolderRequest) (*PaginatedFetchFolderResponse, error) {

	// Fetch all folders by OrgID
	r, err := PaginatedFetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err // Return error if folder retrieval fails
	}

	totalFolders := len(r)
	offset := req.Offset

	// If a token is provided, decode it to get the offset
	if req.NextToken != "" {
		var err error
		offset, err = DecodeToken(req.NextToken)
		if err != nil {
			return nil, err
		}
	}

	// If the offset is beyond the total number of folders, return an empty result.
	if offset >= totalFolders {
		return &PaginatedFetchFolderResponse{
			Folders:   []*Folder{},
			Total:     totalFolders,
			NextToken: "",
		}, nil
	}

	// Determine the end index for the current page of results.
	end := offset + req.Limit
	if end > totalFolders {
		end = totalFolders
	}

	paginatedFolders := r[offset:end]
	newToken := ""
	// If there are more folders, generate a new token for the next page.
	if end < totalFolders {
		newToken = generateNextToken(end)
	}

	// Return the paginated results and the token for the next page.
	return &PaginatedFetchFolderResponse{
		Folders:   paginatedFolders,
		Total:     totalFolders,
		NextToken: newToken,
	}, nil
}

func PaginatedFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}

func generateNextToken(offset int) string {
	return hex.EncodeToString([]byte(fmt.Sprintf("%d", offset)))
}

func DecodeToken(token string) (int, error) {
	data, err := hex.DecodeString(token)
	if err != nil {
		return 0, err
	}
	offset, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return offset, nil
}
