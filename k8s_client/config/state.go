package config

type State int

const (
	Restart State = iota
	Status
	Undo
)
