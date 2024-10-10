package system

import (
	"github.com/t-geindre/golem/pkg/golemutils"
	"testing"
)

func TestAnimation(t *testing.T) {
	a := NewAnimation()
	golemutils.AssertImplementsUpdater("Animation", a, t)
}

func TestCamera(t *testing.T) {
	c := NewCamera()
	golemutils.AssertImplementsUpdater("Camera", c, t)
}

func TestRenderer(t *testing.T) {
	r := NewRenderer(nil)
	golemutils.AssertImplementsDrawer("Renderer", r, t)
}
