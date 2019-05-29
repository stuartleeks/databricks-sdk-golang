package azure

// LibrariesAPI exposes the Libraries API
type LibrariesAPI struct {
	Client DBClient
}

func (a LibrariesAPI) init(client DBClient) LibrariesAPI {
	a.Client = client
	return a
}
