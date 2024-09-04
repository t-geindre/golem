package helper

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

func Damage(l, r golem.Entity, w golem.World) {
	lLife := getLife(l)
	rLife := getLife(r)
	if lLife != nil && rLife != nil {
		lLife.Current, rLife.Current = lLife.Current-rLife.Current, rLife.Current-lLife.Current
		if lLife.Current <= 0 {
			w.RemoveEntity(l)
		}
		if rLife.Current <= 0 {
			w.RemoveEntity(r)
		}
	}
}

func getLife(e golem.Entity) *component.LifeImpl {
	if l, ok := e.(component.Life); ok {
		return l.GetLife()
	}
	return nil
}
