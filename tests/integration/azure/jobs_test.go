package integration

import (
	"encoding/json"
	"os"
	"testing"

	databricks "github.com/xinsnake/databricks-sdk-golang"
	dbAzure "github.com/xinsnake/databricks-sdk-golang/azure"
	dbAzureModels "github.com/xinsnake/databricks-sdk-golang/azure/models"
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
	v, err := json.Marshal(job)
	t.Logf("%+v\n", string(v))
}

func TestList(t *testing.T) {
	c := setup()
	jobs, err := c.Jobs().List()
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	v, err := json.Marshal(jobs)
	t.Logf("%s\n", string(v))
}

func TestRunsList(t *testing.T) {
	c := setup()
	runsView, err := c.Jobs().RunsList(false, false, 0, 0, 0)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	v, err := json.Marshal(runsView)
	t.Logf("%s\n", string(v))
}

func TestRunsSubmit(t *testing.T) {
	c := setup()
	runName := "test-run"
	clusterSpec := dbAzureModels.ClusterSpec{
		NewCluster: &dbAzureModels.NewCluster{
			SparkVersion: "5.2.x-scala2.11",
			NodeTypeID:   "Standard_DS3_v2",
			NumWorkers:   3,
			SparkEnvVars: &dbAzureModels.SparkEnvPair{
				Key:   "PYSPARK_PYTHON",
				Value: "/databricks/python3/bin/python3",
			},
		},
	}
	jobTask := dbAzureModels.JobTask{
		NotebookTask: &dbAzureModels.NotebookTask{
			NotebookPath: "/test-notebook",
		},
	}
	runResponse, err := c.Jobs().RunsSubmit(runName, clusterSpec, jobTask, 0)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	v, err := json.Marshal(runResponse)
	t.Logf("%s\n", string(v))
}
