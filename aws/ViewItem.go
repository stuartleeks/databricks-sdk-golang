package aws

type ViewItem struct {
	Content string   `json:"content,omitempty"`
	Name    string   `json:"name,omitempty"`
	Type    ViewType `json:"type,omitempty"`
}
