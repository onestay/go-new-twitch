package twitch

// Pagination data returned from helix paged requests
type Pagination struct {
	Cursor string `json:"cursor"`
}
