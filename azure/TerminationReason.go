package azure

type TerminationReason struct {
	Code       TerminationCode `json:"code,omitempty"`
	Parameters []ParameterPair `json:"parameters,omitempty"`
}
