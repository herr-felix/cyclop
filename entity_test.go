package cyclop

type dummyEntity struct {
	ID       uint64
	typeName string
	blob     []byte
}

func newDummyEntity(ID uint64, typeName string) *dummyEntity {
	return &dummyEntity{
		ID:       ID,
		typeName: typeName,
		blob:     []byte{},
	}
}

func (e *dummyEntity) GetID() uint64 {
	return e.ID
}

func (e *dummyEntity) GetType() string {
	return e.typeName
}

func (e *dummyEntity) Serialize() []byte {
	return e.blob
}

func (e *dummyEntity) Inspect(inspector *Inspector) {
}
