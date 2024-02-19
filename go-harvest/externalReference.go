package goharvest

type ExternalReference struct {
	ID             int    `json:"id"`
	GroupID        int    `json:"group_id"`
	AccountID      int    `json:"account_id"`
	Permalink      string `json:"permalink"`
	Service        string `json:"service"`
	ServiceIconURL string `json:"service_icon_url"`
}
