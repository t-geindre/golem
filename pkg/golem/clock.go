package golem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type clock struct {
	time    time.Time
	timeSec time.Time
	ticks   int
	elapsed time.Duration
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

	var t time.Time
	if c.ticks == ebiten.TPS() {
		// Fix clock precision
		c.timeSec = c.timeSec.Add(time.Second)
		t = c.timeSec
		c.ticks = 0
	} else {
		t = c.time.Add(time.Second / time.Duration(ebiten.TPS()))
	}

	c.elapsed = t.Sub(c.time)
	c.time = t
}

func (c *clock) Elapsed() time.Duration {
	return c.elapsed
}
