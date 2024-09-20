package helper

import (
	"github.com/t-geindre/golem/examples/scenes/component"
	"github.com/t-geindre/golem/examples/scenes/entity"
	"github.com/t-geindre/golem/examples/scenes/system"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/sfnt"
	"image/color"
	"os"
	"time"
)

const LayerAll = 0

func GetSlides(l golem.LayerID) []golem.Entity {
	slides := make([]golem.Entity, 0)

	for _, es := range [][]golem.Entity{
		{
			entity.NewText(LayerAll, "Slide one", .5, .5, getFontFace(), 12, color.RGBA{R: 23, G: 43, B: 77, A: 0xff}),
		},
		{
			entity.NewText(LayerAll, "Games in GO?!", .5, .1, getFontFace(), 20, colornames.Red),
			entity.NewText(LayerAll, "Ho!", .5, .7, getFontFace(), 30, colornames.Black),
		},
		{
			entity.NewText(LayerAll, "Slide three", .5, .5, getFontFace(), 24, color.RGBA{R: 23, G: 43, B: 77, A: 0xff}),
		},
		{
			entity.NewText(LayerAll, "• List item one", .5, .5, getFontFace(), 20, color.RGBA{R: 23, G: 43, B: 77, A: 0xff}),
			entity.NewText(LayerAll, "• List item two", .5, .7, getFontFace(), 20, color.RGBA{R: 23, G: 43, B: 77, A: 0xff}),
		},
	} {
		slides = append(slides, GetSlide(l, "Slide", TransitionFade, time.Millisecond*150, es...))
	}

	return slides
}

func GetSlide(l golem.LayerID, n string, t component.TransitionFunc, td time.Duration, es ...golem.Entity) golem.Entity {
	se := entity.NewScene(l, n, t, td)
	se.Lifecycle.SetUp = func() {
		se.World.AddLayers(LayerAll)
		se.World.AddEntities(es...)
		se.AddSystem(system.NewRenderer())
	}
	se.Lifecycle.TearDown = func() {
		se.World.Clear()
	}

	return se
}

func getFontFace() *sfnt.Font {
	bts, err := os.ReadFile("assets/TruenoSemibold-Z9yl.otf")
	if err != nil {
		panic(err)
	}

	ft, err := sfnt.Parse(bts)
	if err != nil {
		panic(err)
	}

	return ft
}
