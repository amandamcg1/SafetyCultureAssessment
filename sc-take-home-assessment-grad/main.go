package main

import (
	"fmt"
	"os"

	"github.com/amandamcg1/SafetyCultureAssessment/sc-take-home-assessment-grad/folders/folders"
	"github.com/gofrs/uuid"
)

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Usage: go run main.go [token]")
		return
	}

	orgID := uuid.FromStringOrNil(folders.DefaultOrgID)
	limit := 2
	nextToken := ""
	offset := 0

	if len(os.Args) == 2 {
		nextToken = os.Args[1]
		var err error
		offset, err = folders.de
		if err != nil {
			fmt.Printf("Error decoding token: %v\n", err)
			return
		}
	}

	req := &folders.PaginatedFetchFolderRequest{
		OrgID:     orgID,
		Limit:     limit,
		Offset:    offset,
		NextToken: nextToken,
	}

	res, err := folders.PaginatedGetAllFolders(req)
	if err != nil {
		fmt.Printf("Error fetching folders: %v\n", err)
		return
	}

	fmt.Printf("Folders:\n")
	for i, folder := range res.Folders {
		fmt.Printf("Folder %d: %s\n", offset+i+1, folder.Name)
	}
	fmt.Printf("Token: %s\n", res.NextToken)

	if res.NextToken == "" {
		fmt.Println("End of data.")
	}
}
