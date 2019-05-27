package azure

type NotebookTask struct {
	NotebookPath   string      `json:"notebook_path,omitempty"`
	BaseParameters []ParamPair `json:"base_parameters,omitempty"`
}
