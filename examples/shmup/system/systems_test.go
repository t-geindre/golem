package system

import (
	"github.com/t-geindre/golem/pkg/golemutils"
	"testing"
)

func TestAnimation(t *testing.T) {
	golemutils.AssertImplementsUpdater("Animation", NewAnimation(), t)
}

func TestCollision(t *testing.T) {
	golemutils.AssertImplementsUpdater("Collision", NewCollision([]CollisionRule{}), t)
}

func TestControls(t *testing.T) {
	golemutils.AssertImplementsUpdater("Controls", NewControls(), t)
}

func TestDebug(t *testing.T) {
	d := NewDebug(0)
	golemutils.AssertImplementsDrawer("Debug", d, t)
	golemutils.AssertImplementsUpdaterOnce("Debug", d, t)
}

func TestDespawner(t *testing.T) {
	golemutils.AssertImplementsUpdater("Despawner", NewDespawner(), t)
}

func TestMove(t *testing.T) {
	golemutils.AssertImplementsUpdater("Move", NewMove(), t)
}

func TestRenderer(t *testing.T) {
	golemutils.AssertImplementsDrawer("Renderer", NewRenderer(), t)
}

func TestShoot(t *testing.T) {
	golemutils.AssertImplementsUpdater("Shoot", NewShoot(), t)
}

func TestSpawner(t *testing.T) {
	golemutils.AssertImplementsUpdaterOnce("Spawner", NewSpawner(0, 0, []SpawnFunc{}...), t)
}
