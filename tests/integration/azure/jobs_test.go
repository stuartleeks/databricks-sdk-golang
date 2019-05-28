package integration

import (
	"os"
	"testing"

	databricks "github.com/xinsnake/databricks-sdk-golang"
	dbAzure "github.com/xinsnake/databricks-sdk-golang/azure"
)

func setup() dbAzure.DBClient {
	var c dbAzure.DBClient
	var o databricks.DBClientOption
	o.Host = os.Getenv("DBSDK_HOST")
	o.Token = os.Getenv("DBSDK_TOKEN")
	return c.Init(o)
}

func TestList(t *testing.T) {
	c := setup()
	jobs, err := c.Jobs().List()
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	t.Logf("%+v\n", jobs)
}
