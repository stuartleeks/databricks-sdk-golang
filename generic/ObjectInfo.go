package generic

type ObjectInfo struct {
	ObjectType ObjectType `json:"object_type,omitempty"`
	Path string `json:"path,omitempty"`
	Language Language `json:"language,omitempty"`
}