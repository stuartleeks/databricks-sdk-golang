package generic

type ClusterLibraryStatuses struct {
	ClusterID       string              `json:"cluster_id,omitempty"`
	LibraryStatuses []LibraryFullStatus `json:"library_statuses,omitempty"`
}
