package helper

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
	"math"
)

func TransitionFade(entity golem.Entity, v, _ float64) {
	op := component.GetOpacity(entity)
	if op != nil {
		op.Value = float32(v)
	}
}

func TransitionScale(entity golem.Entity, v, _ float64) {
	scale := component.GetScale(entity)
	if scale != nil {
		scale.Value = v
	}
}

func TransitionHorizontal(entity golem.Entity, v, d float64) {
	pos := component.GetPosition(entity)
	if pos != nil {
		if d > 0 {
			pos.RelX = -1 + v
		} else {
			pos.RelX = 1 - v
		}
	}
}

func TransitionVertical(entity golem.Entity, v, d float64) {
	pos := component.GetPosition(entity)
	if pos != nil {
		if d > 0 {
			pos.RelY = -1 + v
		} else {
			pos.RelY = 1 - v
		}
	}
}

func TransitionRotate(entity golem.Entity, v, d float64) {
	const Deg360ToRad = 360 * math.Pi / 180
	rot := component.GetRotation(entity)
	if rot != nil {
		rot.Angle = d * v * Deg360ToRad
	}
}

func TransitionNone(_ golem.Entity, _, _ float64) {
	// Do nothing
}
