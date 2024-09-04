package component

import "github.com/t-geindre/golem/pkg/golem"

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

func GetLife(e golem.Entity) *LifeImpl {
	if l, ok := e.(Life); ok {
		return l.GetLife()
	}
	return nil
}
