package component

import (
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type TransitionFunc func(entity golem.Entity, v float64)

//go:generate golem component Transition
type Transition struct {
	Transitioning bool
	Start         time.Time
	Duration      time.Duration
	Apply         TransitionFunc
}

func NewTransition(apply TransitionFunc, duration time.Duration) *Transition {
	return &Transition{
		Duration: duration,
		Apply:    apply,
	}
}
