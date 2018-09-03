package cyclop

type dummySearch struct{}

func newDummySearch() *dummySearch {
	return &dummySearch{}
}

func (ds *dummySearch) Index(entityID uint64, result SearchResult, tokens string) error {
	return nil
}

// Search search for results matching the query with a given typeID.
// If offset and limit are 0, it will return every the results.
func (ds *dummySearch) Search(typeID uint16, query string, offset uint32, limit uint32) ([]SearchResult, error) {
	return []SearchResult{}, nil
}

// Remove a result in the search database.
func (ds *dummySearch) Unindex(entityID uint64) error {
	return nil
}
