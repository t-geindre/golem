package component

type Life interface {
	GetLife() *LifeImpl
}

type LifeImpl struct {
	Max, Current int
}

func NewLife(max int) *LifeImpl {
	return &LifeImpl{
		Max:     max,
		Current: max,
	}
}

func (l *LifeImpl) GetLife() *LifeImpl {
	return l
}
