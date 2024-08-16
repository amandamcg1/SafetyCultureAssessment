package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID uuid.UUID
}

type FetchFolderResponse struct {
	Folders []*Folder
}

type PaginatedFetchFolderRequest struct {
	OrgID     uuid.UUID
	Limit     int
	Offset    int
	NextToken string
}

type PaginatedFetchFolderResponse struct {
	Folders   []*Folder
	Total     int
	NextToken string
}
