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

func (c *Collision) Update(left golem.Entity, w golem.World, _ golem.Clock) {
	rules, ok := c.rules[left.GetLayer()]
	if !ok {
		return
	}

	leftCol := component.GetCollider(left)
	if leftCol == nil {
		return
	}

	for _, rule := range rules {
		for _, right := range w.GetEntities(rule.Right) {
			if left == right {
				continue
			}

			rightCol := component.GetCollider(right)
			if rightCol == nil {
				continue
			}

			if c.collides(leftCol, rightCol) {
				rule.Handler(left, right, w)
			}
		}
	}
}

func (c *Collision) collides(l, r *component.Collider) bool {
	return l.Px < r.Px+r.Width &&
		l.Px+l.Width > r.Px &&
		l.Py < r.Py+r.Height &&
		l.Py+l.Height > r.Py
}
