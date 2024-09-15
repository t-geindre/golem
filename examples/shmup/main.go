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
	const (
		LayerBackground = iota
		LayerEnemies
		LayerPlayer
		LayerBullets
		LayerDebug
	)

	ebiten.SetWindowTitle("Golem example - Shmup")
	ebiten.SetWindowSize(600, 800)
	ebiten.SetVsyncEnabled(false)

	g := golemutils.NewGame()

	g.World.AddLayers(LayerBackground, LayerEnemies, LayerPlayer, LayerBullets, LayerDebug)

	collisionRules := []system.CollisionRule{
		{LayerEnemies, LayerBullets, helper.Damage},
		{LayerEnemies, LayerPlayer, helper.Damage},
	}

	g.World.AddSystem(system.NewSpawner(LayerBackground, time.Millisecond*500, entity.NewSparkle))
	g.World.AddSystem(system.NewSpawner(
		LayerEnemies, time.Millisecond*500,
		entity.NewEnemyLips, entity.NewEnemyAllan, entity.NewEnemyBonbon,
	))
	g.World.AddSystem(system.NewControls())
	g.World.AddSystem(system.NewShoot())
	g.World.AddSystem(system.NewMove())
	g.World.AddSystem(system.NewCollision(collisionRules))
	g.World.AddSystem(system.NewDespawner())
	g.World.AddSystem(system.NewAnimation())
	g.World.AddSystem(system.NewRenderer())
	g.World.AddSystem(system.NewDebug(LayerDebug))
	g.World.AddSystem(golemutils.NewMetrics(LayerDebug, time.Millisecond*100))

	g.World.AddEntity(entity.NewPlayer(LayerPlayer, LayerBullets))

	ebiten.RunGame(g)
}
