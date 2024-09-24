package component

import (
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type TransitionFunc func(entity golem.Entity, v, d float64)
type TransitionEaseFunc func(v float64) float64

//go:generate golem component Transition
type Transition struct {
	Transitioning bool
	Start         time.Time
	Duration      time.Duration
	Apply         TransitionFunc
	Ease          TransitionEaseFunc
	Direction     float64 // 1 (next) or -1 (prev)
}

func NewTransition(apply TransitionFunc, ease TransitionEaseFunc, duration time.Duration) *Transition {
	return &Transition{
		Duration: duration,
		Apply:    apply,
		Ease:     ease,
	}
}
