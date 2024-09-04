package component

import "github.com/t-geindre/golem/pkg/golem"

type Collider interface {
	GetCollider() *ColliderImpl
}

type ColliderImpl struct {
	Px, Py, ShiftX, ShiftY, Width, Height float64
}

func NewCollider(sx, sy, w, h float64) *ColliderImpl {
	return &ColliderImpl{
		ShiftX: sx,
		ShiftY: sy,
		Width:  w,
		Height: h,
	}
}

func (c *ColliderImpl) GetCollider() *ColliderImpl {
	return c
}

func GetCollider(e golem.Entity) *ColliderImpl {
	if c, ok := e.(Collider); ok {
		return c.GetCollider()
	}
	return nil
}
