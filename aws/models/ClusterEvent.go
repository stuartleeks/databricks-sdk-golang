package models

type ClusterEvent struct {
	ClusterID string           `json:"cluster_id,omitempty"`
	Timestamp int64            `json:"timestamp,omitempty"`
	Type      ClusterEventType `json:"type,omitempty"`
	Details   AwsAttributes    `json:"details,omitempty"`
}
