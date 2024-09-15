package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/examples/shmup/entity"
	"github.com/t-geindre/golem/examples/shmup/helper"
	"github.com/t-geindre/golem/examples/shmup/system"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

func main() {
	const wWidth, wHeight = 600, 800
	const (
		LayerBackground = iota
		LayerEnemies
		LayerPlayer
		LayerBullets
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Shmup")
	ebiten.SetWindowSize(wWidth, wHeight)
	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()

	g.World.AddLayers(LayerBackground, LayerEnemies, LayerPlayer, LayerBullets, LayerDebug)

	collisionRules := []system.CollisionRule{
		{LayerEnemies, LayerBullets, helper.Damage},
		{LayerEnemies, LayerPlayer, helper.Damage},
	}

	g.World.AddSystem(system.NewSpawner(LayerBackground, 0, wWidth, 0, time.Millisecond*500, entity.NewSparkle))
	g.World.AddSystem(system.NewSpawner(
		LayerEnemies, 20, wWidth-20, 0, time.Millisecond*500,
		entity.NewEnemyLips, entity.NewEnemyAllan, entity.NewEnemyBonbon,
	))
	g.World.AddSystem(system.NewControls())
	g.World.AddSystem(system.NewShoot())
	g.World.AddSystem(system.NewMove())
	g.World.AddSystem(system.NewCollision(collisionRules))
	g.World.AddSystem(system.NewDespawner(0, wHeight, 10))
	g.World.AddSystem(system.NewAnimation())
	g.World.AddSystem(system.NewRenderer())
	g.World.AddSystem(system.NewDebug(LayerDebug))
	g.World.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))

	g.World.AddEntity(entity.NewPlayer(
		LayerPlayer, LayerBullets, wWidth/2, wHeight-50,
		// TODO Constraint needs improvement, this is not flexible at all
		10*helper.Scale, wWidth-10*helper.Scale, 10*helper.Scale, wHeight-10*helper.Scale,
	))

	ebiten.RunGame(g)
}
