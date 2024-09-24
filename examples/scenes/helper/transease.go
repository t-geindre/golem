package helper

// Sources https://easings.net/

import (
	"math"
)

const c5 = (2 * math.Pi) / 4.5

func TransitionEaseLinear(v float64) float64 {
	return v
}

func TransitionEaseSin(v float64) float64 {
	return math.Sin(v * math.Pi / 2)
}

func TransitionEaseQuad(v float64) float64 {
	return 1 - (1-v)*(1-v)
}

func TransitionEaseCubic(v float64) float64 {
	return 1 - math.Pow(1-v, 3)
}

func TransitionEaseBounce(v float64) float64 {
	return 1 - math.Pow(1-v, 2)*(2.70158*(1-v)-1.70158)
}

func TransitionEaseBack(v float64) float64 {
	const c1 = 1.70158
	const c3 = c1 + 1

	return 1 + c3*math.Pow(v-1, 3) + c1*math.Pow(v-1, 2)
}
