package system

import (
	"github.com/t-geindre/golem/pkg/golemutils"
	"testing"
)

func TestBounce(t *testing.T) {
	b := NewBounce()
	golemutils.AssertImplementsUpdaterOnce("Bounce", b, t)
	golemutils.AssertImplementsUpdater("Bounce", b, t)
}

func TestMove(t *testing.T) {
	golemutils.AssertImplementsUpdater("Move", NewMove(), t)
}

func TestRenderer(t *testing.T) {
	golemutils.AssertImplementsDrawer("Renderer", NewRenderer(), t)
}

func TestSpawner(t *testing.T) {
	golemutils.AssertImplementsUpdaterOnce("Spawner", NewSpawner(0, 0), t)
}
