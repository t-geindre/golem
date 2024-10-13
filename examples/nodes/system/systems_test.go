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

func TestGeometry(t *testing.T) {
	golemutils.AssertImplementsUpdater("Geometry", NewGeometry(), t)
}

func TestMove(t *testing.T) {
	golemutils.AssertImplementsUpdater("Move", NewMove(), t)
}

func TestRenderer(t *testing.T) {
	golemutils.AssertImplementsDrawer("Render", NewRenderer(), t)
}
