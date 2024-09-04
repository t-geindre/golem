package system

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
	"image/color"
)

const margin = 5
const cw = 6

type Debug struct {
	l      golem.LayerID
	showHb bool
}

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

	str := "[F1] Show hitboxes"
	if d.showHb {
		str = "[F1] Hide hitboxes"
	}
	ebitenutil.DebugPrintAt(screen, str, screen.Bounds().Dx()-cw*len(str)-margin, margin)
}

func (d *Debug) UpdateOnce(w golem.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		d.showHb = !d.showHb
	}
}

func (d *Debug) Draw(e golem.Entity, screen *ebiten.Image, w golem.World) {
	if !d.showHb {
		return
	}

	hitbox := component.GetCollider(e)
	if hitbox == nil {
		return
	}

	vector.StrokeRect(screen, float32(hitbox.Px), float32(hitbox.Py), float32(hitbox.Width), float32(hitbox.Height), 1, color.RGBA{
		R: 0xff,
		B: 0xff,
		A: 0xff,
	}, false)
}

func (d *Debug) GetLayer() golem.LayerID {
	return d.l
}
