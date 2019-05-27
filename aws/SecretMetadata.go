package aws

type SecretMetadata struct {
	Key                  string `json:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty"`
}
