package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/t-geindre/golem/pkg/golem"
)

type Fullscreen struct {
}

func NewFullscreen() *Fullscreen {
	return &Fullscreen{}
}

func (f *Fullscreen) UpdateOnce(w golem.World) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && ebiten.IsKeyPressed(ebiten.KeyAlt) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
}
