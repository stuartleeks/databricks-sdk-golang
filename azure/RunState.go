package azure

type RunState struct {
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty"`
	ResultState    RunResultState    `json:"result_state,omitempty"`
	StateMessage   string            `json:"state_message,omitempty"`
}
