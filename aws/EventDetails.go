package aws

type EventDetails struct {
	CurrentNumWorkers   int32             `json:"current_num_workers,omitempty"`
	TargetNumWorkers    int32             `json:"target_num_workers,omitempty"`
	PreviousAttributes  ClusterAttributes `json:"previous_attributes,omitempty"`
	Attributes          ClusterAttributes `json:"attributes,omitempty"`
	PreviousClusterSize ClusterSize       `json:"previous_cluster_size,omitempty"`
	ClusterSize         ClusterSize       `json:"cluster_size,omitempty"`
}
