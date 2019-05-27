package databricks

// DataBrickClient is the client to init for communicating with DataBricks, Python source:
// https://github.com/databricks/databricks-cli/blob/master/databricks_cli/sdk/api_client.py
// based on commit: v0.8.6 (cb22ab21327b7d1c85bb45404f7b1115e1fa9dcf)
type DataBrickClient struct {
	User               string
	Password           string
	Host               string
	Token              string
	DefaultHeaders     map[string]string
	InsecureSkipVerify bool
	TimeoutSeconds     int
}
