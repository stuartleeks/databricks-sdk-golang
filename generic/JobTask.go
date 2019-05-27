package generic

type JobTask struct {
	NotebookTask    NotebookTask    `json:"notebook_task,omitempty"`
	SparkJarTask    SparkJarTask    `json:"spark_jar_task,omitempty"`
	SparkPythonTask SparkPythonTask `json:"spark_python_task,omitempty"`
	SparkSubmitTask SparkSubmitTask `json:"spark_submit_task,omitempty"`
}
