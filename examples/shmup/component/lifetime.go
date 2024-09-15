package component

import "time"

//go:generate golem component Lifetime
type Lifetime struct {
	Start time.Time
	Life  time.Duration
}

func NewLifetime(life time.Duration) *Lifetime {
	return &Lifetime{
		Life: life,
	}
}
