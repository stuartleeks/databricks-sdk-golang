package generic

// RunLifeCycleState (The life cycle state of a run.)
type RunLifeCycleState string

const (
	PENDING        = "PENDING"
	RUNNING        = "RUNNING"
	TERMINATING    = "TERMINATING"
	TERMINATED     = "TERMINATED"
	SKIPPED        = "SKIPPED"
	INTERNAL_ERROR = "INTERNAL_ERROR"
)
