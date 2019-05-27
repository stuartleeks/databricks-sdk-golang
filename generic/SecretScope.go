package generic

type SecretScope struct {
	Name string `json:"name,omitempty"`
	BackendType ScopedBackendType `json:"backend_type,omitempty"`
}