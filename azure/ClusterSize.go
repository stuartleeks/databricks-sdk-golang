package azure

type ClusterSize struct {
	NumWorkers int32     `json:"num_workers,omitempty"`
	Autoscale  AutoScale `json:"autoscale,omitempty"`
}
