package aws

type InitScriptInfo struct {
	Dbfs DbfsStorageInfo `json:"dbfs,omitempty"`
	S3   S3StorageInfo   `json:"s3,omitempty"`
}
