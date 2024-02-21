package goharvest

// Documentation is thin, but this appears to be a reference to a celandar
// event. I.e., this would be used to match a time entry with a
// GCal import.
type ExternalReference struct {
	ID             int    `json:"id"`
	GroupID        int    `json:"group_id"`
	AccountID      int    `json:"account_id"`
	Permalink      string `json:"permalink"`
	Service        string `json:"service"`
	ServiceIconURL string `json:"service_icon_url"`
}
