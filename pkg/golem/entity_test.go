package golem

import "testing"

func TestEntityHasWorld(t *testing.T) {
	e := NewEntity(0)

	hw := e.hasWorld()
	if hw {
		t.Errorf("Entity has world")
	}

	hw = e.hasWorld()
	if !hw {
		t.Errorf("Entity has no world")
	}
}
