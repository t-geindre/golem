package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/entity"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/examples/shmup/system"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	const wWidth, wHeight = 600, 800
	const (
		LayerEnemies = iota
		LayerPlayer
		LayerBullets
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Shmup")
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetVsyncEnabled(false)

	w := golem.NewWorld()

	w.AddLayers(LayerEnemies, LayerPlayer, LayerBullets, LayerDebug)

	collisionRules := []system.CollisionRule{
		{LayerEnemies, LayerBullets, helper.Damage},
		{LayerEnemies, LayerPlayer, helper.Damage},
	}

	w.AddSystem(system.NewSpawner(LayerEnemies, 20, wWidth-20, 0, entity.NewEnemy, time.Millisecond*500))
	w.AddSystem(system.NewControls())
	w.AddSystem(system.NewShoot())
	w.AddSystem(system.NewMove())
	w.AddSystem(system.NewConstraint())
	w.AddSystem(system.NewCollision(collisionRules))
	w.AddSystem(system.NewDespawner(0, wHeight, 10))
	w.AddSystem(system.NewAnimation())
	w.AddSystem(system.NewRenderer())
	w.AddSystem(system.NewDebug(LayerDebug))
	w.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))

	w.AddEntity(entity.NewPlayer(
		LayerPlayer, LayerBullets, wWidth/2, wHeight-50,
		// TODO Constraint needs improvement, this is not flexible at all
		10*helper.Scale, wWidth-10*helper.Scale, 10*helper.Scale, wHeight-10*helper.Scale,
	))

	ebiten.RunGame(NewGame(w))
}
