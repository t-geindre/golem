package helper

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"math"
)

const c5 = (2 * math.Pi) / 4.5

func TransitionEaseLinear(v, _ float64) float64 {
	return v
}

func TransitionEaseSin(v, d float64) float64 {
	if component.TransitionIn == d {
		return 1 - math.Cos(v*math.Pi/2)
	}
	return math.Sin(v * math.Pi / 2)
}

func TransitionEaseQuad(v, d float64) float64 {
	if component.TransitionIn == d {
		return v * v
	}
	return 1 - (1-v)*(1-v)
}

func TransitionEaseCubic(v, d float64) float64 {
	if component.TransitionIn == d {
		return v * v * v
	}
	return 1 - math.Pow(1-v, 3)
}

func TransitionEaseBounce(v, d float64) float64 {
	if component.TransitionIn == d {
		return v * v * (2.70158*v - 1.70158)
	}
	return 1 - math.Pow(1-v, 2)*(2.70158*(1-v)-1.70158)
}
