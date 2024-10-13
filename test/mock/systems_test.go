package mock

import (
	"github.com/t-geindre/golem/pkg/golemutils"
	"testing"
)

func TestSysA(t *testing.T) {
	s := NewSysA()
	golemutils.AssertImplementsUpdater("System A", s, t)
	golemutils.AssertImplementsDrawer("System A", s, t)
}

func TestSysB(t *testing.T) {
	s := NewSysB()
	golemutils.AssertImplementsUpdater("System B", s, t)
	golemutils.AssertImplementsDrawer("System B", s, t)
}

func TestSysC(t *testing.T) {
	s := NewSysC()
	golemutils.AssertImplementsUpdater("System C", s, t)
	golemutils.AssertImplementsDrawer("System C", s, t)
}
