package test

import (
	"github.com/t-geindre/golem/test/mock"
	"testing"
)

func TestWorldRemoveEntity(t *testing.T) {
	w := mock.GetSimpleWorld(10)
	es := w.GetEntities(mock.LayerA)[:3]
	l := len(w.GetEntities(mock.LayerA))

	for _, e := range es {
		w.RemoveEntity(e)
		if len(w.GetEntities(mock.LayerA)) != l {
			t.Errorf("Preflush, expected %d entities, got %d", l, len(w.GetEntities(mock.LayerA)))
		}
		l--
		w.Flush()
		if len(w.GetEntities(mock.LayerA)) != l {
			t.Errorf("Postflush, expected %d entities, got %d", l, len(w.GetEntities(mock.LayerA)))
		}
	}

}
