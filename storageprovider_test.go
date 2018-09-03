package cyclop

import "errors"

type dummyStorage struct {
}

func newDummyStorage() *dummyStorage {
	return &dummyStorage{}
}

func (sp *dummyStorage) Store(entityID uint64, entityType uint16, blob []byte) error {
	// Let's say we store something here...

	if string(blob) == "FAIL!" {
		return errors.New("Artificial store error")
	}
	return nil
}
func (sp *dummyStorage) Retreive(entityID uint64, entityType uint16) ([]byte, error) {
	return []byte{}, nil
}
func (sp *dummyStorage) Delete(entityID uint64) error {
	return nil
}
func (sp *dummyStorage) Link(parentID uint64, childID uint64) error {
	return nil
}
func (sp *dummyStorage) Unlink(parentID uint64, childID uint64) error {
	return nil
}
func (sp *dummyStorage) ChildrenOf(entityID uint64) ([]uint64, error) {
	return []uint64{}, nil
}
