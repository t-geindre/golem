package component

import "github.com/t-geindre/golem/pkg/golem"

type DeathSpawnFunc func(l golem.LayerID, x, y float64) golem.Entity

//go:generate golem component Life
type Life struct {
	Max, Current int
	DeathSpawn   DeathSpawnFunc
}

func NewLife(max int, dsf DeathSpawnFunc) *Life {
	return &Life{
		Max:        max,
		Current:    max,
		DeathSpawn: dsf,
	}
}
