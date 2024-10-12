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
	golemutils.AssertImplementsUpdaterOnce("Renderer", NewRenderer(0, 0), t)
}

func TestSpriteRenderer(t *testing.T) {
	golemutils.AssertImplementsDrawer("SpriteRenderer", NewSpriteRenderer(0, 0), t)
}

func TestTextRenderer(t *testing.T) {
	golemutils.AssertImplementsDrawer("TextRenderer", NewTextRenderer(0, 0), t)
}

func TestScene(t *testing.T) {
	s := NewScene(0, []golem.Entity{}...)
	golemutils.AssertImplementsUpdater("Scene", s, t)
	golemutils.AssertImplementsUpdaterOnce("Scene", s, t)
}
