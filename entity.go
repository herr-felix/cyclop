package cyclop

// Entity ...
type Entity interface {
	GetID() uint64
	GetType() string
	GetBlob() []byte
	Inspect(*Inspector) []Warning
}
