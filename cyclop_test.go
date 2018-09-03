package cyclop

import "testing"

func dummyCyclop() *Cyclop {
	return NewCyclop(newDummyStorage(), newDummySearch(), newDummyCache())
}

func TestRegisterType(t *testing.T) {
	const overusedTypeName = "pako"

	c := dummyCyclop()

	err := c.RegisterType(overusedTypeName)

	if err != nil {
		t.Error(err)
	}

	if !c.isTypeRegistered(overusedTypeName) {
		t.Errorf("Type '%s' should be registered", overusedTypeName)
	}

	// Register the same type
	err = c.RegisterType(overusedTypeName)
	if err == nil {
		t.Errorf("Registering the same type should return an error")
	}

	// While we are at it, we check if the types "getters" works well
	ID := c.getTypeIDByName(overusedTypeName)
	if ID != 0 {
		t.Errorf("The only type doesn't have an ID = 0")
	}

	name := c.getTypeNameByID(ID)
	if name != overusedTypeName {
		t.Errorf("Could not find a type with the ID '0'")
	}
}

func TestSave(t *testing.T) {
	const testTypeName = "test"
	c := dummyCyclop()

	e := newDummyEntity(1, testTypeName)

	err := c.Save(e)
	if err == nil {
		t.Errorf("Should not save an non-registered type")
	}

	c.RegisterType(testTypeName)

	err = c.Save(e)
	if err != nil {
		t.Error(err)
	}

	e.data = dummyData{shouldFail: true}
	err = c.Save(e)
	if err == nil {
		t.Errorf("Storage should have failled")
	}
}
