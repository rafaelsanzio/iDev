package models

// GetInfo model
type GetInfo struct {
	Enterprise  string `json:"enterprise"`
	About       string `json:"about"`
	ProjectLink string `json:"projectLink"`
	Creator     string `json:"creator"`
}
