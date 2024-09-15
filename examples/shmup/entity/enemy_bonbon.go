package entity

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type EnemyBonbon struct {
	*Enemy
}

func NewEnemyBonbon(l golem.LayerID) golem.Entity {
	return &EnemyBonbon{
		Enemy: NewEnemy(l,
			component.NewFrame(helper.Assets["en_bonbon_f1"], time.Millisecond*500),
			component.NewFrame(helper.Assets["en_bonbon_f2"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_bonbon_f3"], time.Millisecond*50),
			component.NewFrame(helper.Assets["en_bonbon_f4"], time.Millisecond*100),
		),
	}
}
