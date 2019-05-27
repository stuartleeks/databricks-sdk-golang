package azure

type ClusterSpec struct {
	ExistingClusterID string     `json:"existing_cluster_id,omitempty"`
	NewCluster        NewCluster `json:"new_cluster,omitempty"`
	Libraries         []Library  `json:"libraries,omitempty"`
}
