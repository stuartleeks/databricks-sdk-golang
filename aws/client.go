package aws

import databricks "github.com/xinsnake/databricks-sdk-golang"

// DBClient is the client for Azure implements DBClient
type DBClient struct {
	Option databricks.DBClientOption
}

// Init initializes the client
func (c DBClient) Init(option databricks.DBClientOption) DBClient {
	c.Option = option
	return c
}

// Clusters returns an instance of ClustersAPI
func (c DBClient) Clusters() ClustersAPI {
	var clustersAPI ClustersAPI
	return clustersAPI.init(c)
}

// Dbfs returns an instance of DbfsAPI
func (c DBClient) Dbfs() DbfsAPI {
	var dbfsAPI DbfsAPI
	return dbfsAPI.init(c)
}

// Jobs returns an instance of JobsAPI
func (c DBClient) Jobs() JobsAPI {
	var jobsAPI JobsAPI
	return jobsAPI.init(c)
}

func (c *DBClient) performQuery(method, path string, data interface{}, headers map[string]string) ([]byte, error) {
	return databricks.PerformQuery(c.Option, method, path, data, headers)
}
