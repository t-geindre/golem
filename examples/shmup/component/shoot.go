package component

import (
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type BulletSpawn func(id golem.LayerID) golem.Entity

type Shoot interface {
	GetShoot() *ShootImpl
}

type ShootImpl struct {
	Shooting bool
	Rate     time.Duration
	Last     time.Time
	Spawn    BulletSpawn
	AtX, AtY float64
	Layer    golem.LayerID
}

func NewShoot(rate time.Duration, x, y float64, spawn BulletSpawn, l golem.LayerID) *ShootImpl {
	return &ShootImpl{
		Shooting: false,
		Rate:     rate,
		Last:     time.Now(),
		Spawn:    spawn,
		AtX:      x,
		AtY:      y,
		Layer:    l,
	}
}

func (s *ShootImpl) GetShoot() *ShootImpl {
	return s
}

func GetShoot(e golem.Entity) *ShootImpl {
	if s, ok := e.(Shoot); ok {
		return s.GetShoot()
	}
	return nil
}
