package generic

// RunLifeCycleState (The life cycle state of a run.)
type RunLifeCycleState string

const (
	RunLifeCycleStatePending       = "PENDING"
	RunLifeCycleStateRunning       = "RUNNING"
	RunLifeCycleStateTerminating   = "TERMINATING"
	RunLifeCycleStateTerminated    = "TERMINATED"
	RunLifeCycleStateSkipped       = "SKIPPED"
	RunLifeCycleStateInternalError = "INTERNAL_ERROR"
)
