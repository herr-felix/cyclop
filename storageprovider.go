package cyclop

// StorageProvider ...
type StorageProvider interface {
	Store(entityID uint64, entityType uint16, blob []byte) error
	Retreive(entityID uint64, entityType uint16) ([]byte, error)
	Delete(entityID uint64) error
	Link(parentID uint64, childID uint64) error
	Unlink(parentID uint64, childID uint64) error
	ChildrenOf(entityID uint64) ([]uint64, error)
}
