package golem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
	Tick()
}

type clock struct {
	time time.Time
}

func newClock() Clock {
	return &clock{}
}

func (c *clock) Now() time.Time {
	return c.time
}

func (c *clock) Since(t time.Time) time.Duration {
	return c.time.Sub(t)
}

func (c *clock) Tick() {
	c.time = c.time.Add(time.Second / time.Duration(ebiten.TPS()))
}
