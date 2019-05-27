package databricks

import (
	"github.com/xinsnake/databricks-sdk-golang/aws"
	"github.com/xinsnake/databricks-sdk-golang/azure"
)

// AwsDBClient is the client for Azure
type AwsDBClient struct {
	Option DBClientOption

	Clusters         aws.ClustersAPI
	Dbfs             aws.DbfsAPI
	Groups           aws.GroupsAPI
	InstanceProfiles aws.InstanceProfilesAPI
	Jobs             aws.JobsAPI
	Libraries        aws.LibrariesAPI
	Scim             aws.ScimAPI
	Secrets          aws.SecretsAPI
	Token            aws.TokenAPI
	Workspace        aws.WorkspaceAPI
}

// PerformQuery provies an abstraction for performQuery
func (c *AwsDBClient) PerformQuery(
	method, path string, data map[string]interface{}, headers map[string]string) ([]byte, error) {
	return performQuery(c.Option, method, path, data, headers)
}

// AzureDBClient is the client for Azure
type AzureDBClient struct {
	Option DBClientOption

	Clusters  azure.ClustersAPI
	Dbfs      azure.DbfsAPI
	Groups    azure.GroupsAPI
	Jobs      azure.JobsAPI
	Libraries azure.LibrariesAPI
	Scim      azure.ScimAPI
	Secrets   azure.SecretsAPI
	Token     azure.TokenAPI
	Workspace azure.WorkspaceAPI
}

// PerformQuery provies an abstraction for performQuery
func (c *AzureDBClient) PerformQuery(
	method, path string, data map[string]interface{}, headers map[string]string) ([]byte, error) {
	return performQuery(c.Option, method, path, data, headers)
}
