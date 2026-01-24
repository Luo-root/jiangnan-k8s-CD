package model

type StatusCode int

const (
	Success StatusCode = iota
	ParameterFail
	ApplyFail
	RolloutFail
)
