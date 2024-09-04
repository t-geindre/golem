package component

import "github.com/hajimehoshi/ebiten/v2"

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
