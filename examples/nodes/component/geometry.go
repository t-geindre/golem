package component

import "github.com/hajimehoshi/ebiten/v2"

type Geometry struct {
	*ebiten.GeoM
}

//go:generate golem component Geometry
func NewGeometry() *Geometry {
	return &Geometry{GeoM: &ebiten.GeoM{}}
}
