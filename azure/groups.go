package azure

// GroupsAPI exposes the Groups API
type GroupsAPI struct {
	Client DBClient
}

func (a GroupsAPI) init(client DBClient) GroupsAPI {
	a.Client = client
	return a
}
