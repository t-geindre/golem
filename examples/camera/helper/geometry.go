package helper

import "image"

func RectMulF(r image.Rectangle, f float64) image.Rectangle {
	return image.Rect(
		int(float64(r.Min.X)*f),
		int(float64(r.Min.Y)*f),
		int(float64(r.Max.X)*f),
		int(float64(r.Max.Y)*f),
	)
}

func PointMulF(p image.Point, f float64) image.Point {
	return image.Point{
		X: int(float64(p.X) * f),
		Y: int(float64(p.Y) * f),
	}
}
