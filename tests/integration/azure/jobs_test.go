package integration

import (
	"encoding/json"
	"os"
	"testing"
	"time"

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

func TestRuns(t *testing.T) {
	c := setup()
	runName := "test-run"
	clusterSpec := dbAzureModels.ClusterSpec{
		NewCluster: &dbAzureModels.NewCluster{
			SparkVersion: "5.2.x-scala2.11",
			NodeTypeID:   "Standard_DS3_v2",
			NumWorkers:   3,
			SparkEnvVars: map[string]string{
				"PYSPARK_PYTHON": "/databricks/python3/bin/python3",
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
	v, _ := json.Marshal(runResponse)
	t.Logf("%s\n", string(v))

	time.Sleep(5 * time.Second)

	err = c.Jobs().RunsCancel(runResponse.RunID)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	time.Sleep(5 * time.Second)
	err = c.Jobs().RunsDelete(runResponse.RunID)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
}
