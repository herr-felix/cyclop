package cyclop

import "time"

// InfractionType ...
type InfractionType uint8

const (
	// InfractionInformation ...
	InfractionInformation InfractionType = iota
	// InfractionWarning ...
	InfractionWarning InfractionType = iota
	// InfractionError ...
	InfractionError InfractionType = iota
	// InfractionCritical ...
	InfractionCritical InfractionType = iota
)

// Infraction ...
type Infraction struct {
	EntityID    uint64
	Level       uint8
	Desctiption string
	Timestamp   time.Time
	Snoozed     bool
}
