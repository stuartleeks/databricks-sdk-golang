package azure

import "github.com/xinsnake/databricks-sdk-golang/azure/models"

// JobsAPI exposes Jobs API endpoints
type JobsAPI struct {
}

// Create creates a new job
func (a *JobsAPI) Create(settings models.JobSettings) (models.Job, error) {

}

// List lists all jobs
func (a *JobsAPI) List() ([]models.Job, error) {

}

// Delete deletes a job by ID
func (a *JobsAPI) Delete(jobID int64) (error) {

}

// Get gets a job by ID
func (a *JobsAPI) Get(jobID int64) (models.Job, error) {

}

func (a *JobsAPI) Reset() {

}

func (a *JobsAPI) RunNow() {

}

func (a *JobsAPI) RunsSubmit() {

}

func (a *JobsAPI) RunsList() {

}

func (a *JobsAPI) RunsGet() {

}

func (a *JobsAPI) RunsExport() {

}

func (a *JobsAPI) RunsCancel() {

}

func (a *JobsAPI) RunsGetOutput() {

}

func (a *JobsAPI) RunsDelete() {

}
