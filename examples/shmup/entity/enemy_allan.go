package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type EnemyAllan struct {
	*Enemy
}

func NewEnemyAllan(l golem.LayerID) golem.Entity {
	return &EnemyAllan{
		Enemy: NewEnemy(l,
			component.NewFrame(helper.Assets["en_allan_f1"], time.Millisecond*500),
			component.NewFrame(helper.Assets["en_allan_f2"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_allan_f3"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_allan_f4"], time.Millisecond*100),
			component.NewFrame(helper.Assets["en_allan_f5"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_allan_f6"], time.Millisecond*50),
		),
	}
}
