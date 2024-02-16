package goharvest

type Pagination struct {
	PerPage      int  `json:"per_page"`
	TotalPages   int  `json:"total_pages"`
	TotalEntries int  `json:"total_entries"`
	NextPage     *int `json:"next_page"`
	PreviousPage *int `json:"previous_page"`
	Page         int  `json:"page"`
	Links        struct {
		First    string  `json:"first"`
		Next     *string `json:"next"`
		Previous *string `json:"previous"`
		Last     string  `json:"last"`
	} `json:"links"`
}
