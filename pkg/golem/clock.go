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
	time    time.Time
	timeSec time.Time
	ticks   int
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
	c.ticks++

	// Fix clock precision
	if c.ticks == ebiten.TPS() {
		c.timeSec = c.timeSec.Add(time.Second)
		c.time = c.timeSec
		c.ticks = 0
		return
	}

	c.time = c.time.Add(time.Second / time.Duration(ebiten.TPS()))
}
