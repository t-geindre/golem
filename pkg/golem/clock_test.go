package golem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reflect"
	"testing"
	"time"
)

func TestClockNow(t *testing.T) {
	type fields struct {
		tickRate int
		ticks    int
	}
	tests := []struct {
		name string
		args fields
		want time.Time
	}{
		{
			name: "Now() after 1 ticks at 60 TPS",
			args: fields{tickRate: 60, ticks: 1},
			want: time.Time{}.Add(time.Second / time.Duration(60)),
		},
		{
			name: "Now() after 10 ticks at 60 TPS",
			args: fields{tickRate: 60, ticks: 10},
			want: time.Time{}.Add((time.Second / time.Duration(60)) * 10),
		},
		{
			name: "Now() after 60 ticks at 60 TPS",
			args: fields{tickRate: 60, ticks: 60},
			want: time.Time{}.Add(time.Second),
		},
		{
			name: "Now() after 210 ticks at 70 TPS",
			args: fields{tickRate: 70, ticks: 210},
			want: time.Time{}.Add(time.Second * 3),
		},
		{
			name: "Now() after 190 ticks at 60 TPS",
			args: fields{tickRate: 60, ticks: 190},
			want: time.Time{}.Add(time.Second*3 + (time.Second/time.Duration(60))*10),
		},
		{
			name: "Now() after 1 ticks at 120 TPS",
			args: fields{tickRate: 120, ticks: 1},
			want: time.Time{}.Add(time.Second / time.Duration(120)),
		},
		{
			name: "Now() after 10 ticks at 120 TPS",
			args: fields{tickRate: 120, ticks: 10},
			want: time.Time{}.Add((time.Second / time.Duration(120)) * 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ebiten.SetTPS(tt.args.tickRate)
			c := newClock()

			for i := 0; i < tt.args.ticks; i++ {
				c.Tick()
			}

			if got := c.Now(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClockSince(t *testing.T) {
	type fields struct {
		tickRate int
		ticks    int
		from     time.Time
	}
	tests := []struct {
		name string
		args fields
		want time.Duration
	}{
		{
			name: "Since(0), after 1 tick at 60 TPS",
			args: fields{tickRate: 60, ticks: 1, from: time.Time{}},
			want: time.Second / time.Duration(60),
		},
		{
			name: "Since(0), after 10 ticks at 10 TPS",
			args: fields{tickRate: 10, ticks: 10, from: time.Time{}},
			want: time.Second,
		},
		{
			name: "Since(500ms), after 10 ticks at 10 TPS",
			args: fields{tickRate: 10, ticks: 10, from: time.Time{}.Add(500 * time.Millisecond)},
			want: time.Millisecond * 500,
		},
		{
			name: "Since(1s), after 10 ticks at 10 TPS",
			args: fields{tickRate: 10, ticks: 10, from: time.Time{}.Add(time.Second)},
			want: time.Duration(0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ebiten.SetTPS(tt.args.tickRate)
			c := newClock()

			for i := 0; i < tt.args.ticks; i++ {
				c.Tick()
			}

			if got := c.Since(tt.args.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClockElapsed(t *testing.T) {
	type fields struct {
		ticks int
		tps   int
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name:   "Elapsed after 1 tick at 60 TPS",
			fields: fields{ticks: 1, tps: 60},
			want:   time.Second / time.Duration(60),
		},
		{
			name:   "Elapsed after 100 tick at 60 TPS",
			fields: fields{ticks: 1, tps: 60},
			want:   time.Second / time.Duration(60),
		},
		{
			name:   "Elapsed after 100 tick at 120 TPS",
			fields: fields{ticks: 100, tps: 120},
			want:   time.Second / time.Duration(120),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ebiten.SetTPS(tt.fields.tps)
			c := newClock()
			for i := 0; i < tt.fields.ticks; i++ {
				c.Tick()
			}
			if got := c.Elapsed(); got != tt.want {
				t.Errorf("Elapsed() = %v, want %v", got, tt.want)
			}
		})
	}
}
