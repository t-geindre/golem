package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
)

type Sprite interface {
	GetSprite() *SpriteImpl
}

func NewSprite(img *ebiten.Image) *SpriteImpl {
	return &SpriteImpl{Img: img, Opt: &ebiten.DrawImageOptions{}}
}

type SpriteImpl struct {
	Img *ebiten.Image
	Opt *ebiten.DrawImageOptions
}

func (s *SpriteImpl) GetSprite() *SpriteImpl {
	return s
}

func GetSprite(e golem.Entity) *SpriteImpl {
	if s, ok := e.(Sprite); ok {
		return s.GetSprite()
	}
	return nil
}
