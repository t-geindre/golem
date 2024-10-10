package system

import (
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"testing"
)

func TestAnimation(t *testing.T) {
	golemutils.AssertImplementsUpdater("Animation", NewAnimation(), t)
}

func TestRenderer(t *testing.T) {
	golemutils.AssertImplementsDrawer("Renderer", NewRenderer(), t)
}

func TestScene(t *testing.T) {
	s := NewScene(0, []golem.Entity{}...)
	golemutils.AssertImplementsUpdater("Scene", s, t)
	golemutils.AssertImplementsUpdaterOnce("Scene", s, t)
}
