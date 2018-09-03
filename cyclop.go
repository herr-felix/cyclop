package cyclop

import "errors"

// Cyclop controls everything...
type Cyclop struct {
	storage     StorageProvider
	search      SearchProvider
	cache       CacheProvider
	typesIDName map[uint16]string
	typesNameID map[string]uint16
}

// NewCyclop create a new Cyclop instance
func NewCyclop(storage StorageProvider, search SearchProvider, cache CacheProvider) *Cyclop {
	return &Cyclop{
		storage:     storage,
		search:      seach,
		cache:       cache,
		typesNameID: make(map[string]uint16),
		typesIDName: make(map[uint16]string),
	}
}

func (c *Cyclop) getTypeIDByName(name string) uint16 {
	return c.typesNameID[name]
}

// RegisterType ...
func (c *Cyclop) RegisterType(typeName string) error {

	if _, exists := c.typesNameID[Name]; exists {
		return errors.New("Type already exists")
	}

	ID := len(c.typesIDName)

	c.typesIDName[ID] = typeName

	return nil
}

// Save save an entity
func (c *Cyclop) Save(entity Entity) error {

	err := c.storage.Store(entity.GetID(), entity.GetType(), entity.Serialize())

	if err != nil {
		return err
	}

	return nil
}
