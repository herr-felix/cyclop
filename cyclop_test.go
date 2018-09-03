package cyclop

import "testing"

func DummyCyclop() *Cyclop {
	return NewCyclop(newDummyStorage(), newDummySearch(), newDummyCache())
}

func TestRegisterType(t *testing.T) {

	c := DummyCyclop()

	c.PANIC

}
