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
	ebiten.SetFullscreen(true)
	return f
}

func (f *Fullscreen) UpdateOnce(_ golem.World, _ golem.Clock) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		ebiten.SetFullscreen(false)
	}
}
