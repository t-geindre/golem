package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t-geindre/golem/pkg/golem"
)

type Fullscreen struct {
}

func NewFullscreen() *Fullscreen {
	f := &Fullscreen{}
	//f.setFullscreen(true)
	return f
}

func (f *Fullscreen) UpdateOnce(w golem.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt) {
		f.setFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		f.setFullscreen(false)
	}
}

func (f *Fullscreen) setFullscreen(fs bool) {
	ebiten.SetFullscreen(fs)
	if fs {
		ebiten.SetCursorMode(ebiten.CursorModeHidden)
		return
	}
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
}
