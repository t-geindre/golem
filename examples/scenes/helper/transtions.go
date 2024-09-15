package helper

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/pkg/golem"
)

func TransitionFade(entity golem.Entity, v float64) {
	op := component.GetOpacity(entity)
	if op != nil {
		op.Value = float32(v)
	}
}

func TransitionScale(entity golem.Entity, v float64) {
	scale := component.GetScale(entity)
	if scale != nil {
		scale.Value = v
	}
}

func TransitionNone(_ golem.Entity, _ float64) {
	// Do nothing
}
