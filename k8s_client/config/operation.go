package config

type Operation int

const (
	Apply Operation = iota
	Delete
	Rollout
	Get
)
