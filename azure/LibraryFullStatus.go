package azure

type LibraryFullStatus struct {
	Library                 Library              `json:"library,omitempty"`
	Status                  LibraryInstallStatus `json:"status,omitempty"`
	Messages                []string             `json:"messages,omitempty"`
	IsLibraryForAllClusters bool                 `json:"is_library_for_all_clusters,omitempty"`
}
