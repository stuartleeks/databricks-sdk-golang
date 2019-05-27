package aws

type JobSettings struct {
	ExistingClusterID      string                `json:"existing_cluster_id,omitempty"`
	NewCluster             NewCluster            `json:"new_cluster,omitempty"`
	NotebookTask           NotebookTask          `json:"notebook_task,omitempty"`
	SparkJarTask           SparkJarTask          `json:"spark_jar_task,omitempty"`
	SparkPythonTask        SparkPythonTask       `json:"spark_python_task,omitempty"`
	SparkSubmitTask        SparkSubmitTask       `json:"spark_submit_task,omitempty"`
	Name                   string                `json:"name,omitempty"`
	Libraries              []Library             `json:"libraries,omitempty"`
	EmailNotifications     JobEmailNotifications `json:"email_notifications,omitempty"`
	TimeoutSeconds         int32                 `json:"timeout_seconds,omitempty"`
	MaxRetries             int32                 `json:"max_retries,omitempty"`
	MinRetryIntervalMillis int32                 `json:"min_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool                  `json:"retry_on_timeout,omitempty"`
	Schedule               CronSchedule          `json:"schedule,omitempty"`
	MaxConcurrentRuns      int32                 `json:"max_concurrent_runs,omitempty"`
}
