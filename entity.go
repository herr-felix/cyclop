package cyclop

// Entity ...
type Entity interface {
	GetID() uint64
	GetType() string
	Load(interface{}) error
	Freeze() (interface{}, error)
	Inspect(*Inspector)
}
