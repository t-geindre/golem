package system

import (
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
)

type CollisionRule struct {
	Left, Right golem.LayerID
	Handler     func(l, r golem.Entity, w golem.World)
}

type Collision struct {
	rules map[golem.LayerID][]CollisionRule
}

func NewCollision(rules []CollisionRule) *Collision {
	c := &Collision{
		rules: make(map[golem.LayerID][]CollisionRule),
	}
	for _, rule := range rules {
		c.rules[rule.Left] = append(c.rules[rule.Left], rule)
	}
	return c
}

func (c *Collision) Update(left golem.Entity, w golem.World) {
	rules, ok := c.rules[left.GetLayer()]
	if !ok {
		return
	}

	leftCol := c.getCollider(left)
	if leftCol == nil {
		return
	}

	for _, rule := range rules {
		for _, right := range w.GetEntities(rule.Right) {
			if left == right {
				continue
			}

			rightCol := c.getCollider(right)
			if rightCol == nil {
				continue
			}

			if leftCol.CollidesWith(rightCol) {
				rule.Handler(left, right, w)
			}
		}
	}
}

func (c *Collision) getCollider(e golem.Entity) *component.ColliderImpl {
	col, ok := e.(component.Collider)
	if !ok {
		return nil
	}
	return col.GetCollider()
}
