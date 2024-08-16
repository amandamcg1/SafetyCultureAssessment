package main

import (
	"fmt"
	"os"

	"github.com/amandamcg1/SafetyCultureAssessment/sc-take-home-assessment-grad/folders"
	"github.com/gofrs/uuid"
)

func main() {

	// Check if more than one argument is provided; if so, print usage instructions and exit.
	if len(os.Args) > 2 {
		fmt.Println("Usage: go run main.go [token]")
		return
	}

	orgID := uuid.FromStringOrNil(folders.DefaultOrgID)
	limit := 10 // Change Limit for more or less folders
	nextToken := ""
	offset := 0

	// If a token is provided as a command-line argument, decode it to get the offset.
	if len(os.Args) == 2 {
		nextToken = os.Args[1]
		var err error
		offset, err = folders.DecodeToken(nextToken)
		if err != nil {
			fmt.Printf("Error decoding token: %v\n", err)
			return
		}
	}

	// Create a request for paginated folder fetching with the specified parameters.
	req := &folders.PaginatedFetchFolderRequest{
		OrgID:     orgID,
		Limit:     limit,
		Offset:    offset,
		NextToken: nextToken,
	}

	// Fetch the paginated folders using the request parameters.
	res, err := folders.PaginatedGetAllFolders(req)
	if err != nil {
		fmt.Printf("Error fetching folders: %v\n", err)
		return
	}

	// Print the list of folders.
	fmt.Printf("Folders:\n")
	for i, folder := range res.Folders {
		fmt.Printf("Folder %d: %s\n", offset+i+1, folder.Name)
	}

	// Print the next token if there are more folders to fetch.
	// Print a message indicating the end of data if no next token is provided.
	if res.NextToken != "" {
		fmt.Printf("Token: %s\n", res.NextToken)
	} else {
		fmt.Println("End of data.")
	}
}
