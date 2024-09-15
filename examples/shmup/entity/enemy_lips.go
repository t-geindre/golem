package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type EnemyLips struct {
	*Enemy
}

func NewEnemyLips(l golem.LayerID, px, py float64) golem.Entity {
	return &EnemyLips{
		Enemy: NewEnemy(l, px, py,
			component.NewFrame(helper.Assets["en_lips_f1"], time.Millisecond*500),
			component.NewFrame(helper.Assets["en_lips_f2"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_lips_f3"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_lips_f4"], time.Millisecond*200),
			component.NewFrame(helper.Assets["en_lips_f5"], time.Millisecond*50),
		),
	}
}
