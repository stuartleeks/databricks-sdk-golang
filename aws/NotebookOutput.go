package aws

type NotebookOutput struct {
	Result    string `json:"result,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`
}
