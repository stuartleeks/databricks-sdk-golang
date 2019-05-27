package generic

type ClusterCloudProviderNodeInfo struct {
	Status             ClusterCloudProviderNodeStatus `json:"status,omitempty"`
	AvailableCoreQuota int32                          `json:"available_core_quota,omitempty"`
	TotalCoreQuota     int32                          `json:"total_core_quota,omitempty"`
}
