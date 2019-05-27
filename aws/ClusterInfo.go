package aws

type ClusterInfo struct {
	NumWorkers             int32            `json:"num_workers,omitempty"`
	AutoScale              AutoScale        `json:"autoscale,omitempty"`
	ClusterID              string           `json:"cluster_id,omitempty"`
	CreatorUserName        string           `json:"creator_user_name,omitempty"`
	Driver                 SparkNode        `json:"driver,omitempty"`
	Executors              []SparkNode      `json:"executors,omitempty"`
	SparkContextID         int64            `json:"spark_context_id,omitempty"`
	JdbcPort               int32            `json:"jdbc_port,omitempty"`
	ClusterName            string           `json:"cluster_name,omitempty"`
	SparkVersion           string           `json:"spark_version,omitempty"`
	SparkConf              SparkConfPair    `json:"spark_conf,omitempty"`
	AwsAttributes          AwsAttributes    `json:"aws_attributes,omitempty"`
	NodeTypeID             string           `json:"node_type_id,omitempty"`
	DriverNodeTypeID       string           `json:"driver_node_type_id,omitempty"`
	SSHPublicKeys          []string         `json:"ssh_public_keys,omitempty"`
	CustomTags             []ClusterTag     `json:"custom_tags,omitempty"`
	ClusterLogConf         ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	InitScripts            []InitScriptInfo `json:"init_scripts,omitempty"`
	SparkEnvVars           SparkEnvPair     `json:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32            `json:"autotermination_minutes,omitempty"`
	EnableElasticDisk      bool             `json:"enable_elastic_disk,omitempty"`
	ClusterSource          AwsAvailability  `json:"cluster_source,omitempty"`
	State                  ClusterState     `json:"state,omitempty"`
	StateMessage           string           `json:"state_message,omitempty"`
	StartTime              int64            `json:"start_time,omitempty"`
	TerminateTime          int64            `json:"terminate_time,omitempty"`
	LastStateLossTime      int64            `json:"last_state_loss_time,omitempty"`
	LastActivityTime       int64            `json:"last_activity_time,omitempty"`
	ClusterMemoryMb        int64            `json:"cluster_memory_mb,omitempty"`
	ClusterCores           float32          `json:"cluster_cores,omitempty"`
	DefaultTags            []ClusterTag     `json:"default_tags,omitempty"`
	ClusterLogStatus       LogSyncStatus    `json:"cluster_log_status,omitempty"`
	TerminationReason      S3StorageInfo    `json:"termination_reason,omitempty"`
}
