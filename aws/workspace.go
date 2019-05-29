package aws

// WorkspaceAPI exposes the Workspace API
type WorkspaceAPI struct {
	Client DBClient
}

func (a WorkspaceAPI) init(client DBClient) WorkspaceAPI {
	a.Client = client
	return a
}

// Delete
// Export
// GetStatus
// Import
// List
// Mkdirs
