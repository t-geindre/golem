package golemutils

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"runtime"
	"time"
)

type Metrics struct {
	l golem.LayerID
	*Panel
}

func NewMetrics(l golem.LayerID, rate time.Duration) *Metrics {
	return &Metrics{
		l:     l,
		Panel: NewPanel(l, metricsRefresh, rate),
	}
}

func metricsRefresh(w golem.World) string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return fmt.Sprintf(
		"FPS: %.2f TPS: %.2f\nSys Memory: %v MB Entities: %d",
		ebiten.ActualFPS(),
		ebiten.ActualTPS(),
		m.Sys/1024/1024,
		w.Size(),
	)
}
