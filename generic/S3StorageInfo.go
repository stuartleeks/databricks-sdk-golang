package generic

type S3StorageInfo struct {
	Destination      string `json:"destination,omitempty"`
	Region           string `json:"region,omitempty"`
	Endpoint         string `json:"endpoint,omitempty"`
	EnableEncryption bool   `json:"enable_encryption,omitempty"`
	EncryptionType   string `json:"encryption_type,omitempty"`
	KmsKey           string `json:"kms_key,omitempty"`
	CannedACL        string `json:"canned_acl,omitempty"`
}
