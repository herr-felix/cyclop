package cyclop

const (
	WarningLevelInformation   = 0
	WarningLevelWarning       = 1
	WarningLevelError         = 2
	WarningLevelCriticalError = 3
)

// Warning is and Entity warning.
// Warnings can have multiple levels of severity.
type Warning struct {
	Entity      uint64
	Level       uint8
	Desctiption string
	Snoozed     bool
}
