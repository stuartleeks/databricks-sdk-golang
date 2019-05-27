package generic

type RunResultState string

const (
	SUCCESS  = "SUCCESS"
	FAILED   = "FAILED"
	TIMEDOUT = "TIMEDOUT"
	CANCELED = "CANCELED"
)
