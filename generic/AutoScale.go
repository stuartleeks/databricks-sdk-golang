package generic

// AutoScale (No description)
type AutoScale struct {
	MinWorkers int32 `json:"min_workers,omitempty"`
	MaxWorkers int32 `json:"max_workers,omitempty"`
}
