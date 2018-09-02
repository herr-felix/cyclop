package cyclop

import "time"

// CacheProvider ...
type CacheProvider interface {
	// SetDefaultTTL sets default time to live.
	// If negative, default TTL will be set to infinity.
	SetDefaultTTL(ttl time.Duration) error

	// Exists verify if an entity is in cache.
	Exists(entityID uint64) (bool, error)

	// Get retreive an entity by its ID.
	// Return nil, CacheErrorNotFound if nothing is found.
	Get(entityID uint64) (*RawEntity, error)

	// Put place an entity in cache with a TTL.
	// If the TTL is 0, default will be used. If negative, no TTL is set.
	// If the entity already exists in cache, it will be overwritten.
	Put(entity *RawEntity, TTL time.Duration) error

	// Delete simply remove an entity from the cache.
	Delete(entityID uint64) error
}
