package golemutils

import (
	"github.com/t-geindre/golem/pkg/golem"
	"testing"
)

func AssertImplementsDrawer(name string, s any, t *testing.T) {
	_, ok := s.(golem.Drawer)
	if !ok {
		t.Errorf(`"%s" system must implement golem.Drawer interface`, name)
	}
}

func AssertImplementsUpdater(name string, s any, t *testing.T) {
	_, ok := s.(golem.Updater)
	if !ok {
		t.Errorf(`"%s" system must implement golem.Updater interface`, name)
	}
}

func AssertImplementsDrawerOnce(name string, s any, t *testing.T) {
	_, ok := s.(golem.DrawerOnce)
	if !ok {
		t.Errorf(`"%s" system must implement golem.DrawerOnce interface`, name)
	}
}

func AssertImplementsUpdaterOnce(name string, s any, t *testing.T) {
	_, ok := s.(golem.UpdaterOnce)
	if !ok {
		t.Errorf(`"%s" system must implement golem.UpdaterOnce interface`, name)
	}
}
