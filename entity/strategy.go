package entity

type strategy int

const (
	DistStrtegyEQUAL strategy = iota + 1
	DistStrtegyExact
	DistStrtegyPercent
)
