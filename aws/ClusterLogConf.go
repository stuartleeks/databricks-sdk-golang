package aws

type ClusterLogConf struct {
	Dbfs DbfsStorageInfo `json:"dbfs,omitempty"`
	S3   S3StorageInfo   `json:"s3,omitempty"`
}
