package cyclop

// SearchResult ...
type SearchResult struct {
	EntityID uint64
	TypeID   uint16
	// MsgPack blob
	Blob []byte
}

// SearchProvider ...
type SearchProvider interface {
	// IndexEntity add a search record to the search database.
	// If there already is a record with the same entityID, it will be updated.
	Index(entityID uint64, result SearchResult, tokens string) error

	// Search search for results matching the query with a given typeID.
	// If offset and limit are 0, it will return every the results.
	Search(typeID uint16, query string, offset uint32, limit uint32) ([]SearchResult, error)

	// Remove a result in the search database.
	Unindex(entityID uint64) error
}
