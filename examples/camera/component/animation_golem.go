// Code generated by Golem. DO NOT EDIT.
// Source: https://github.com/t-geindre/golem

package component

import "github.com/t-geindre/golem/pkg/golem"

type AnimationGolemI interface {
	GetAnimation() *Animation
}

func (p *Animation) GetAnimation() *Animation {
	return p
}

func GetAnimation(e golem.Entity) *Animation {
	if p, ok := e.(AnimationGolemI); ok {
		return p.GetAnimation()
	}
	return nil
}