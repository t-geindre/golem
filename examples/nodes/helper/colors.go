package helper

import (
	"image/color"
	"math/rand"
)

func RandomColor(alpha uint8) color.Color {
	return color.RGBA{
		R: uint8(100 + rand.Intn(155)),
		G: uint8(100 + rand.Intn(155)),
		B: uint8(100 + rand.Intn(155)),
		A: alpha,
	}
}
