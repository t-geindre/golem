package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"time"
)

type Debug struct {
	l golem.LayerID
	*golemutils.Panel
}

const (
	margin = 5
	cw     = 6
	ch     = 16
)

func NewDebug(l golem.LayerID) *Debug {
	d := &Debug{l: l}
	p := golemutils.NewPanel(l, d.refresh, time.Millisecond*100, golemutils.StickTopLeft)
	d.Panel = p
	return d
}

func (d *Debug) refresh(w golem.World) string {
	return fmt.Sprintf(
		"FPS: %.2f TPS: %.2f\nEntities: %d",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
		w.Size(),
	)
}
