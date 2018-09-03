package cyclop

// Entity ...
type Entity interface {
	GetID() uint64
	GetType() string
	Serialize() []byte
	Inspect(*Inspector) []Infraction
}
