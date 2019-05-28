package integration

import (
	"os"
	"testing"

	databricks "github.com/xinsnake/databricks-sdk-golang"
	dbAzure "github.com/xinsnake/databricks-sdk-golang/azure"
)

func setup() dbAzure.DBClient {
	var o databricks.DBClientOption
	o.Host = os.Getenv("DATABRICKS_HOST")
	o.Token = os.Getenv("DATABRICKS_TOKEN")
	var c dbAzure.DBClient
	return c.Init(o)
}

func TestGet(t *testing.T) {
	c := setup()
	job, err := c.Jobs().Get(1)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	t.Logf("%+v\n", job)
}

func TestList(t *testing.T) {
	c := setup()
	jobs, err := c.Jobs().List()
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	t.Logf("%+v\n", jobs)
}

func TestRunsList(t *testing.T) {
	c := setup()
	runsView, err := c.Jobs().RunsList(false, false, 0, 0, 0)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	t.Logf("%+v\n", runsView)
}
