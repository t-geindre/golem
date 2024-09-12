// Code generated by Golem. DO NOT EDIT.
// Source: https://github.com/t-geindre/golem

package component

import "github.com/t-geindre/golem/pkg/golem"

type ColliderGolemI interface {
	GetCollider() *Collider
}

func (p *Collider) GetCollider() *Collider {
	return p
}

func GetCollider(e golem.Entity) *Collider {
	if p, ok := e.(ColliderGolemI); ok {
		return p.GetCollider()
	}
	return nil
}