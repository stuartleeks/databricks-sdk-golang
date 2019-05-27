package models

type NewCluster struct {
	NumWorkers        int32            `json:"num_workers,omitempty"`
	Autoscale         AutoScale        `json:"autoscale,omitempty"`
	ClusterName       string           `json:"cluster_name,omitempty"`
	SparkVersion      string           `json:"spark_version,omitempty"`
	SparkConf         SparkConfPair    `json:"spark_conf,omitempty"`
	AwsAttributes     AwsAttributes    `json:"aws_attributes,omitempty"`
	NodeTypeID        string           `json:"node_type_id,omitempty"`
	DriverNodeTypeID  string           `json:"driver_node_type_id,omitempty"`
	SSHPublicKeys     []string         `json:"ssh_public_keys,omitempty"`
	CustomTags        []ClusterTag     `json:"custom_tags,omitempty"`
	ClusterLogConf    ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	InitScripts       []InitScriptInfo `json:"init_scripts,omitempty"`
	SparkEnvVars      SparkEnvPair     `json:"spark_env_vars,omitempty"`
	EnableElasticDisk bool             `json:"enable_elastic_disk,omitempty"`
}
