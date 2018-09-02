package cyclop

// RawEntity ...
type RawEntity struct {
	ID   uint64
	Type uint16
	Blob []byte
}

// StorageProvider ...
type StorageProvider interface {
	Retreive(entityID uint64) (*RawEntity, error)
	Store(entity *RawEntity) error
	Delete(entityID uint64) error
	Link(parentID uint64, childID uint64) error
	UnLink(parentID uint64, childID uint64) error
	ChildrenOf(entityID uint64) ([]uint64, error)
}
