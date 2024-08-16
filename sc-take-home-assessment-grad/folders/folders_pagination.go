package folders

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/gofrs/uuid"
)

func PaginatedGetAllFolders(req *PaginatedFetchFolderRequest) (*PaginatedFetchFolderResponse, error) {

	// Fetch all folders by OrgID
	r, err := PaginatedFetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	totalFolders := len(r)
	offset := req.Offset

	if req.NextToken != "" {
		var err error
		offset, err = decodeToken(req.NextToken)
		if err != nil {
			return nil, err
		}
	}

	if offset >= totalFolders {
		return &PaginatedFetchFolderResponse{
			Folders:   []*Folder{},
			Total:     totalFolders,
			NextToken: "",
		}, nil
	}

	end := offset + req.Limit
	if end > totalFolders {
		end = totalFolders
	}

	paginatedFolders := r[offset:end]
	newToken := ""
	if end < totalFolders {
		newToken = generateNextToken(end)
	}

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
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", offset)))
}

func decodeToken(token string) (int, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}
	offset, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return offset, nil
}
