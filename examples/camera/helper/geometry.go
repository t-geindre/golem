package helper

import (
	"image"
	"math"
)

func RectMulF(r image.Rectangle, f float64) image.Rectangle {
	return image.Rect(
		int(float64(r.Min.X)*f),
		int(float64(r.Min.Y)*f),
		int(float64(r.Max.X)*f),
		int(float64(r.Max.Y)*f),
	)
}

func ISignCeilFloor(x float64) int {
	if x < 0 {
		return int(math.Floor(x))
	}
	return int(math.Ceil(x))
}
