package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/t-geindre/golem/examples/shmup/component"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/pkg/golemutils"
	"image/color"
)

const margin = 5
const cw = 6

type Debug struct {
	l      golem.LayerID
	showHb bool
	*golemutils.Panel
}

func NewDebug(l golem.LayerID) *Debug {
	d := &Debug{l: l}
	d.Panel = golemutils.NewPanel(l, d.GetDisplay, golemutils.RefreshOnce)
	d.Panel.Stick = golemutils.StickTopRight
	return d
}

func (d *Debug) UpdateOnce(w golem.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		d.showHb = !d.showHb
		d.Panel.Refresh(w)
	}
	d.Panel.UpdateOnce(w)
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
		G: 0xff,
		R: 0xff,
		A: 0xff,
	}, false)
}

func (d *Debug) GetLayer() golem.LayerID {
	return d.l
}

func (d *Debug) GetDisplay(w golem.World) string {
	if d.showHb {
		return "[F1] Hide hitboxes"
	}
	return "[F1] Show hitboxes"
}
