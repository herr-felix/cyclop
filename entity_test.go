package cyclop

import "errors"

type dummyEntity struct {
	ID       uint64
	typeName string
	data     dummyData
}

type dummyData struct {
	shouldFail bool
}

func newDummyEntity(ID uint64, typeName string) *dummyEntity {
	return &dummyEntity{
		ID:       ID,
		typeName: typeName,
		data:     dummyData{shouldFail: false},
	}
}

func (e *dummyEntity) GetID() uint64 {
	return e.ID
}

func (e *dummyEntity) GetType() string {
	return e.typeName
}

func (e *dummyEntity) Freeze() (interface{}, error) {
	if e.data.shouldFail {
		return nil, errors.New("Failling on puspose")
	}
	return e.data, nil
}

func (e *dummyEntity) Load(obj interface{}) (err error) {
	data, ok := obj.(dummyData)
	if !ok {
		return errors.New("Type could not be infered")
	}
	e.data = data
	return nil
}

func (e *dummyEntity) Inspect(inspector *Inspector) {
}
