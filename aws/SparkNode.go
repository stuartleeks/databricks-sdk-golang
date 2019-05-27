package aws

type SparkNode struct {
	PrivateIP         string                 `json:"private_ip,omitempty"`
	PublicDNS         string                 `json:"public_dns,omitempty"`
	NodeID            string                 `json:"node_id,omitempty"`
	InstanceID        string                 `json:"instance_id,omitempty"`
	StartTimestamp    int64                  `json:"start_timestamp,omitempty"`
	NodeAwsAttributes SparkNodeAwsAttributes `json:"node_aws_attributes,omitempty"`
	HostPrivateIP     string                 `json:"host_private_ip,omitempty"`
}
