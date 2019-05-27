package azure

type SecretScope struct {
	Name        string           `json:"name,omitempty"`
	BackendType ScopeBackendType `json:"backend_type,omitempty"`
}
