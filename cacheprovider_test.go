package cyclop

import "time"

type dummyCache struct{}

func newDummyCache() *dummyCache {
	return &dummyCache{}
}

func (cp *dummyCache) SetDefaultTTL(ttl time.Duration) error {
	return nil
}

func (cp *dummyCache) Exists(entityID uint64) (bool, error) {
	return false, nil
}

func (cp *dummyCache) Get(entityID uint64, entityType string) ([]byte, error) {
	return []byte{}, nil
}

func (cp *dummyCache) Put(entityID uint64, entityType uint16, blob []byte, TTL time.Duration) error {
	return nil
}

func (cp *dummyCache) Delete(entityID uint64) error {
	return nil
}
