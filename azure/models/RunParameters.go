package models

type RunParameters struct {
	JarParams         []string    `json:"jar_params,omitempty"`
	NotebookParams    []ParamPair `json:"notebook_params,omitempty"`
	PythonParams      []string    `json:"python_params,omitempty"`
	SparkSubmitParams []string    `json:"spark_submit_params,omitempty"`
}
