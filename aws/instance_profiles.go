package aws

// InstanceProfilesAPI exposes the Instance Profiles API
type InstanceProfilesAPI struct {
	Client DBClient
}

func (a InstanceProfilesAPI) init(client DBClient) InstanceProfilesAPI {
	a.Client = client
	return a
}

// ProfilesAdd
// ProfilesList
// ProfilesRemove