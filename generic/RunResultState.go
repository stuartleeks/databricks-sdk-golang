package generic

type RunResultState string

const (
	RunResultStateSuccess  = "SUCCESS"
	RunResultStateFailed   = "FAILED"
	RunResultStateTimedout = "TIMEDOUT"
	RunResultStateCanceled = "CANCELED"
)
