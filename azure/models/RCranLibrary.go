package models

type RCranLibrary struct {
	Package string `json:"package,omitempty"`
	Repo    string `json:"repo,omitempty"`
}
