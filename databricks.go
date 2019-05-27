package databricks

// DBClient is the client
type DBClient struct {
	Option   DBClientOption
	Clusters ClustersAPI
	Dbfs     DbfsAPI
	Jobs     JobsAPI
}
