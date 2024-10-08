// Code generated by Golem. DO NOT EDIT.
// Source: https://github.com/t-geindre/golem

package component

import "github.com/t-geindre/golem/pkg/golem"

type VelocityGolemI interface {
	GetVelocity() *Velocity
}

func (p *Velocity) GetVelocity() *Velocity {
	return p
}

func GetVelocity(e golem.Entity) *Velocity {
	if p, ok := e.(VelocityGolemI); ok {
		return p.GetVelocity()
	}
	return nil
}