package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/t-geindre/golem/pkg/golem"
)

type Debug struct {
	l golem.LayerID
}

const margin = 5

func NewDebug(l golem.LayerID) *Debug {
	return &Debug{l: l}
}

func (d *Debug) DrawOnce(screen *ebiten.Image, w golem.World) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf(
		"FPS %.0f\nTPS %d\nEntities %d\n",
		ebiten.ActualFPS(),
		ebiten.TPS(),
		w.Size(),
	), margin, margin)
}

func (d *Debug) GetLayer() golem.LayerID {
	return d.l
}
