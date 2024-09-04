package component

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

func (c *ColliderImpl) CollidesWith(other *ColliderImpl) bool {
	return c.Px < other.Px+other.Width &&
		c.Px+c.Width > other.Px &&
		c.Py < other.Py+other.Height &&
		c.Py+c.Height > other.Py
}
