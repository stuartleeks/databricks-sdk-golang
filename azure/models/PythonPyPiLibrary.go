package models

type PythonPyPiLibrary struct {
	Package string `json:"package,omitempty"`
	Repo    string `json:"repo,omitempty"`
}
