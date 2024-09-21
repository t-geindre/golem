package helper

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"image/color"
	"math/rand"
	"os"
	"time"
)

const LayerAll = 0

func GetSlides(l golem.LayerID) []golem.Entity {
	slides := make([]golem.Entity, 0)

	for _, es := range [][]golem.Entity{
		{
			entity.NewText(LayerAll, "Games with GO", .5, .3, .5, .5, getFontFace(200), color.RGBA{R: 0x17, G: 0x2b, B: 0x4d, A: 0xff}),
			entity.NewGopher(LayerAll, component.NewFrame(Assets[0], time.Second*3), component.NewFrame(Assets[1], time.Millisecond*200)),
			entity.NewText(LayerAll, "Why not?", .5, .65, .5, .5, getFontFace(100), color.RGBA{R: 0x49, G: 0x90, B: 0xf9, A: 0xff}),
		},
		{
			entity.NewText(LayerAll, "Suitable for games?", .5, .4, .5, .5, getFontFace(150), color.RGBA{R: 0x17, G: 0x2b, B: 0x4d, A: 0xff}),
			entity.NewText(LayerAll, "It... depends!", .5, .6, .5, .5, getFontFace(100), color.RGBA{R: 0x49, G: 0x90, B: 0xf9, A: 0xff}),
		},
		{
			entity.NewText(LayerAll, "What do you need?", .5, .2, .5, .5, getFontFace(150), color.RGBA{R: 0x17, G: 0x2b, B: 0x4d, A: 0xff}),
			entity.NewText(LayerAll, "• Performances", .3, .5, 0, .5, getFontFace(100), color.RGBA{R: 0x49, G: 0x90, B: 0xf9, A: 0xff}),
			entity.NewText(LayerAll, "• Ecosystem", .3, .65, 0, .5, getFontFace(100), color.RGBA{R: 0x49, G: 0x90, B: 0xf9, A: 0xff}),
		},
		{
			entity.NewText(LayerAll, "Performances", .5, .4, .5, .5, getFontFace(150), color.RGBA{R: 0x17, G: 0x2b, B: 0x4d, A: 0xff}),
			entity.NewText(LayerAll, "GO vs C++ vs Rust VS C#", .5, .6, .5, .5, getFontFace(100), color.RGBA{R: 0x49, G: 0x90, B: 0xf9, A: 0xff}),
		},
	} {
		trans := []component.TransitionFunc{
			TransitionFade,
			TransitionScale,
		}[rand.Intn(2)]
		slides = append(slides, GetSlide(l, "Slide", trans, time.Millisecond*200, es...))
	}

	return slides
}

func GetSlide(l golem.LayerID, n string, t component.TransitionFunc, td time.Duration, es ...golem.Entity) golem.Entity {
	se := entity.NewScene(l, n, t, td)
	se.Lifecycle.SetUp = func() {
		se.World.AddLayers(LayerAll)
		se.World.AddEntities(es...)
		se.AddSystems(
			system.NewSpriteRenderer(),
			system.NewTextRenderer(),
			system.NewAnimation(),
		)
	}
	se.Lifecycle.TearDown = func() {
		se.World.Clear()
	}

	return se
}

func getFontFace(size float64) text.Face {
	bts, err := os.ReadFile("assets/TruenoSemibold-Z9yl.otf")
	if err != nil {
		panic(err)
	}

	ft, err := sfnt.Parse(bts)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(ft, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}

	return text.NewGoXFace(face)
}
