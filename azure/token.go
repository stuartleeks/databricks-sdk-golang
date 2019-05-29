package azure

// TokenAPI exposes the Token API
type TokenAPI struct {
	Client DBClient
}

func (a TokenAPI) init(client DBClient) TokenAPI {
	a.Client = client
	return a
}
