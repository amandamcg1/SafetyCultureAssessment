package main

import (
	"fmt"

	"github.com/amandamcg1/SafetyCultureAssessment/sc-take-home-assessment-grad/folders/folders"
	"github.com/gofrs/uuid"
)

func main() {
	req := &folders.FetchFolderRequest{
		OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	}

	res, err := folders.GetAllFolders(req)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	folders.PrettyPrint(res)
}
