package golemutils

import (
	"github.com/t-geindre/golem/pkg/golem"
	"testing"
)

func TestPanelInterfaces(t *testing.T) {
	var p any = &Panel{}
	t.Run("Panel{} implements golem.DrawOnce", func(t *testing.T) {
		_, ok := p.(golem.DrawerOnce)
		if !ok {
			t.Errorf("Panel{} does not implement golem.DrawerOnce")
		}
	})
	t.Run("Panel{} implements golem.DrawOnce", func(t *testing.T) {
		_, ok := p.(golem.UpdaterOnce)
		if !ok {
			t.Errorf("Panel{} does not implement golem.UpdaterOnce")
		}
	})
}
