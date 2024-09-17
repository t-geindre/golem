package test

import (
	"github.com/t-geindre/golem/test/mock"
	"testing"
)

func TestWorlRemoveEntity(t *testing.T) {
	w := mock.GetSimpleWorld(10)
	e := w.GetEntities(mock.LayerA)[5]

	w.RemoveEntity(e)
	w.Flush()

	if len(w.GetEntities(mock.LayerA)) != 9 {
		t.Errorf("Expected 9 entities, got %d", len(w.GetEntities(mock.LayerA)))
	}
}
