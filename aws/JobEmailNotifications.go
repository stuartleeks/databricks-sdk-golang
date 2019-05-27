package aws

type JobEmailNotifications struct {
	OnStart               []string `json:"on_start,omitempty"`
	OnSuccess             []string `json:"on_success,omitempty"`
	OnFailure             []string `json:"on_failure,omitempty"`
	NoAlertForSkippedRuns bool     `json:"no_alert_for_skipped_runs,omitempty"`
}
