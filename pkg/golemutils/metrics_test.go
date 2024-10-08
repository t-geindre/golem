package golemutils

import (
	"github.com/t-geindre/golem/pkg/golem"
	"testing"
)

func TestMetricsInterfaces(t *testing.T) {
	var m any = &Metrics{}
	t.Run("Metrics{} implements golem.DrawOnce", func(t *testing.T) {
		_, ok := m.(golem.DrawerOnce)
		if !ok {
			t.Errorf("Metrics{} does not implement golem.DrawerOnce")
		}
	})
}
