package helper

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

func Damage(l, r golem.Entity, w golem.World) {
	lLife := component.GetLife(l)
	rLife := component.GetLife(r)

	if lLife == nil || rLife == nil {
		return
	}

	lLife.Current, rLife.Current = lLife.Current-rLife.Current, rLife.Current-lLife.Current
}
