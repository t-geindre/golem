package golemutils

// Mainly inspired by: https://github.com/sedyh/mizu/blob/main/examples/bunnymark/system/metrics.go

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/t-geindre/golem/pkg/golem"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"strings"
	"time"
)

const (
	StickTopLeft = iota
	StickTopRight
	StickBottomLeft
	StickBottomRight
)

type Panel struct {
	Str        string
	Margin     float32
	Padding    float32
	Background color.RGBA
	Foreground color.RGBA
	Ticker     *time.Ticker
	Rate       time.Duration
	Refresh    func(w golem.World) string
	Stick      int
	Layer      golem.LayerID
	Font       font.Face

	width    int
	height   int
	baseline int
}

func NewPanel(layer golem.LayerID, refresh func(w golem.World) string, rate time.Duration, stick int) *Panel {
	p := &Panel{
		Str:        "",
		Margin:     10,
		Padding:    10,
		Background: color.RGBA{R: 0, G: 0, B: 0, A: 0xaa},
		Foreground: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		Rate:       rate,
		Refresh:    refresh,
		Stick:      stick,
		Layer:      layer,
		Font:       basicfont.Face7x13,
	}

	if rate > 0 {
		p.Ticker = time.NewTicker(rate)
	}

	return p
}

func (p *Panel) UpdateOnce(w golem.World) {
	if p.Rate > 0 {
		select {
		case <-p.Ticker.C:
			p.refresh(w)
		default:
		}
		return
	}

	p.refresh(w)
}

func (p *Panel) DrawOnce(screen *ebiten.Image, w golem.World) {
	if len(p.Str) == 0 {
		return
	}

	vector.DrawFilledRect(
		screen,
		p.Margin, p.Margin,
		float32(p.width)+p.Padding*2,
		float32(p.height)+p.Padding*2,
		p.Background,
		false,
	)
	text.Draw(
		screen,
		p.Str,
		basicfont.Face7x13,
		int(p.Padding)+int(p.Margin),
		p.baseline+int(p.Padding)+int(p.Margin),
		colornames.White,
	)
}

func (p *Panel) GetLayer() golem.LayerID {
	return p.Layer
}

func (p *Panel) refresh(w golem.World) {
	p.Str = p.Refresh(w)
	m := p.Font.Metrics()

	width := 0
	height := 0
	for _, part := range strings.Split(p.Str, "\n") {
		w := font.MeasureString(p.Font, part).Round()
		if w > width {
			width = w
		}
		height += m.Height.Round()
	}

	p.width = width
	p.height = height
	p.baseline = m.Ascent.Round()
}
