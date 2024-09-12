package component

import (
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type BulletSpawn func(id golem.LayerID) golem.Entity

//go:generate golem shoot.go
type Shoot struct {
	Shooting bool
	Rate     time.Duration
	Last     time.Time
	Spawn    BulletSpawn
	AtX, AtY float64
	Layer    golem.LayerID
}

func NewShoot(rate time.Duration, x, y float64, spawn BulletSpawn, l golem.LayerID) *Shoot {
	return &Shoot{
		Shooting: false,
		Rate:     rate,
		Last:     time.Now(),
		Spawn:    spawn,
		AtX:      x,
		AtY:      y,
		Layer:    l,
	}
}
