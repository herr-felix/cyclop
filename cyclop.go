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
		search:      search,
		cache:       cache,
		typesNameID: make(map[string]uint16),
		typesIDName: make(map[uint16]string),
	}
}

func (c *Cyclop) getTypeIDByName(name string) uint16 {
	return c.typesNameID[name]
}

func (c *Cyclop) getTypeNameByID(ID uint16) string {
	return c.typesIDName[ID]
}

func (c *Cyclop) isTypeRegistered(name string) (exists bool) {
	_, exists = c.typesNameID[name]
	return exists
}

func (c *Cyclop) registerTypeName(name string) {
	ID := uint16(len(c.typesIDName))

	c.typesIDName[ID] = name
	c.typesNameID[name] = ID
}

// RegisterType ...
func (c *Cyclop) RegisterType(typeName string) error {

	if c.isTypeRegistered(typeName) {
		return errors.New("Type has already been registered")
	}

	c.registerTypeName(typeName)

	return nil
}

// Save save an entity
func (c *Cyclop) Save(entity Entity) error {

	// Get the ID of the entity's type
	if !c.isTypeRegistered(entity.GetType()) {
		return errors.New("Type is not registered")
	}

	err := c.storage.Store(entity.GetID(), c.getTypeIDByName(entity.GetType()), entity.Serialize())

	if err != nil {
		return err
	}

	return nil
}
