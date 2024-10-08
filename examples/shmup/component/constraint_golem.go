// Code generated by Golem. DO NOT EDIT.
// Source: https://github.com/t-geindre/golem

package component

import "github.com/t-geindre/golem/pkg/golem"

type ConstraintGolemI interface {
	GetConstraint() *Constraint
}

func (p *Constraint) GetConstraint() *Constraint {
	return p
}

func GetConstraint(e golem.Entity) *Constraint {
	if p, ok := e.(ConstraintGolemI); ok {
		return p.GetConstraint()
	}
	return nil
}