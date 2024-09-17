package golemutils

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"time"
)

type Metrics struct {
	l golem.LayerID
	*Panel
	rx, ry int
}

func NewMetrics(l golem.LayerID, rate time.Duration) *Metrics {
	m := &Metrics{l: l}
	m.Panel = NewPanel(l, m.metricsRefresh, rate)
	return m
}

func (m *Metrics) DrawOnce(screen *ebiten.Image, w golem.World) {
	m.rx, m.ry = screen.Bounds().Dx(), screen.Bounds().Dy()
	m.Panel.DrawOnce(screen, w)
}

func (m *Metrics) GetLayer() golem.LayerID {
	return m.l
}

func (m *Metrics) metricsRefresh(w golem.World) string {
	return fmt.Sprintf(
		"FPS: %.2f TPS: %.2f\nRes: %dx%d Entities: %d",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
		m.rx, m.ry,
		w.Size(),
	)
}
