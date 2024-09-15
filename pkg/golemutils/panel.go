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
	"sync"
	"time"
)

const (
	StickTopLeft = iota
	StickTopRight
	StickBottomLeft
	StickBottomRight
	RefreshNever = -1
	RefreshOnce  = -2
)

type Panel struct {
	Str        string
	Margin     float32
	Padding    float32
	Background color.RGBA
	Foreground color.RGBA
	Ticker     *time.Ticker
	Rate       time.Duration
	GetStr     func(w golem.World) string
	Stick      int
	Layer      golem.LayerID
	Font       font.Face

	once     sync.Once
	width    float32
	height   float32
	baseline int
}

func NewPanel(layer golem.LayerID, refresh func(w golem.World) string, rate time.Duration) *Panel {
	// Todo a panel should be an entity updated and drawn by a unique system
	p := &Panel{
		Str:        "",
		Margin:     10,
		Padding:    10,
		Background: color.RGBA{R: 0, G: 0, B: 0, A: 0xaa},
		Foreground: color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		Rate:       rate,
		GetStr:     refresh,
		Stick:      StickTopLeft,
		Layer:      layer,
		Font:       basicfont.Face7x13,
	}

	if rate > 0 {
		p.Ticker = time.NewTicker(rate)
	}

	return p
}

func (p *Panel) UpdateOnce(w golem.World) {
	if p.Rate == RefreshNever {
		return
	}

	if p.Rate == RefreshOnce {
		p.once.Do(func() {
			p.Refresh(w)
		})
		return
	}

	if p.Rate > 0 {
		select {
		case <-p.Ticker.C:
			p.Refresh(w)
		default:
		}
		return
	}

	p.Refresh(w)
}

func (p *Panel) DrawOnce(screen *ebiten.Image, w golem.World) {
	if len(p.Str) == 0 {
		return
	}

	rx, ry := float32(0), float32(0)
	rw, rh := p.width+p.Padding*2, p.height+p.Padding*2

	switch p.Stick {
	case StickTopLeft:
		rx, ry = p.Margin, p.Margin
	case StickTopRight:
		rx, ry = float32(screen.Bounds().Dx())-rw-p.Margin, p.Margin
	case StickBottomLeft:
		rx, ry = p.Margin, float32(screen.Bounds().Dy())-rh-p.Margin
	case StickBottomRight:
		rx, ry = float32(screen.Bounds().Dx())-rw-p.Margin, float32(screen.Bounds().Dy())-rh-p.Margin
	default:
		panic("Unknown sticking position")
	}

	vector.DrawFilledRect(screen, rx, ry, rw, rh, p.Background, false)
	text.Draw(
		screen,
		p.Str,
		basicfont.Face7x13,
		int(rx)+int(p.Padding),
		p.baseline+int(ry)+int(p.Padding),
		colornames.White,
	)
}

func (p *Panel) GetLayer() golem.LayerID {
	return p.Layer
}

func (p *Panel) Refresh(w golem.World) {
	p.Str = p.GetStr(w)
	m := p.Font.Metrics()

	width := float32(0)
	height := float32(0)
	for _, part := range strings.Split(p.Str, "\n") {
		w := float32(font.MeasureString(p.Font, part).Round())
		if w > width {
			width = w
		}
		height += float32(m.Height.Round())
	}

	p.width = width
	p.height = height
	p.baseline = m.Ascent.Round()
}
