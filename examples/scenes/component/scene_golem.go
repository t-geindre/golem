// Code generated by Golem. DO NOT EDIT.
// Source: https://github.com/t-geindre/golem

package component

import "github.com/t-geindre/golem/pkg/golem"

type SceneGolemI interface {
	GetScene() *Scene
}

func (p *Scene) GetScene() *Scene {
	return p
}

func GetScene(e golem.Entity) *Scene {
	if p, ok := e.(SceneGolemI); ok {
		return p.GetScene()
	}
	return nil
}