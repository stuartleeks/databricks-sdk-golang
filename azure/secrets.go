package azure

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	Client DBClient
}

func (a SecretsAPI) init(client DBClient) SecretsAPI {
	a.Client = client
	return a
}
