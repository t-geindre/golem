package component

import (
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type TransitionFunc func(entity golem.Entity, v, d float64)
type TransitionEaseFunc func(v, d float64) float64

const (
	TransitionIn   float64 = 1
	TransitionOut  float64 = -1
	TransitionNext float64 = 1
	TransitionPrev float64 = -1
)

//go:generate golem component Transition
type Transition struct {
	Transitioning bool
	Start         time.Time
	Duration      time.Duration
	Apply         TransitionFunc
	Ease          TransitionEaseFunc
	Direction     float64 // 1 (next) or -1 (prev) todo fix implementation between transition and easing
	InOut         float64 // 1 (in) or -1 (out) todo they both share the same state, should be separated
}

func NewTransition(apply TransitionFunc, ease TransitionEaseFunc, duration time.Duration) *Transition {
	return &Transition{
		Duration: duration,
		Apply:    apply,
		Ease:     ease,
	}
}
