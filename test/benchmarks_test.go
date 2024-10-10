package test

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/test/mock"
	"testing"
)

const (
	Size100K = 1e5
	Size1M   = 1e6
)

func BenchmarkWorldUpdate(b *testing.B) {
	type args struct {
		world golem.World
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Update simple world with 100K entities",
			args: args{world: mock.GetSimpleWorld(Size100K)},
		},
		{
			name: "Update simple world with 1M entities",
			args: args{world: mock.GetSimpleWorld(Size1M)},
		},
		{
			name: "Update complex world with 100K entities",
			args: args{world: mock.GetComplexWorld(Size100K)},
		},
		{
			name: "Update complex world with 1M entities",
			args: args{world: mock.GetComplexWorld(Size1M)},
		},
		{
			name: "Update dead world with 100K entities",
			args: args{world: mock.GetDeadWorld(Size100K)},
		},
		{
			name: "Update dead world with 1M entities",
			args: args{world: mock.GetDeadWorld(Size1M)},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.args.world.Update()
			}
		})
	}
}

func BenchmarkWorldDraw(b *testing.B) {
	type args struct {
		world golem.World
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Draw simple world with 100K entities",
			args: args{world: mock.GetSimpleWorld(Size100K)},
		},
		{
			name: "Draw simple world with 1M entities",
			args: args{world: mock.GetSimpleWorld(Size1M)},
		},
		{
			name: "Draw complex world with 100K entities",
			args: args{world: mock.GetComplexWorld(Size100K)},
		},
		{
			name: "Draw complex world with 1M entities",
			args: args{world: mock.GetComplexWorld(Size1M)},
		},
		{
			name: "Draw dead world with 100K entities",
			args: args{world: mock.GetDeadWorld(Size100K)},
		},
		{
			name: "Draw dead world with 1M entities",
			args: args{world: mock.GetDeadWorld(Size1M)},
		},
	}

	img := &ebiten.Image{}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.args.world.Draw(img)
			}
		})
	}
}

func BenchmarkWorldAddRemoveSingleEntity(b *testing.B) {
	w := mock.GetDeadWorld(Size1M)
	es := w.GetEntities(mock.LayerA)
	e := es[len(es)/2]
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.RemoveEntity(e)
		w.Flush()

		w.AddEntity(e)
		w.Flush()
	}
}

func BenchmarkWorldAddRemoveMultipleEntities(b *testing.B) {
	w := mock.GetDeadWorld(Size1M)
	es := w.GetEntities(mock.LayerA)
	ea := es[0]
	eb := es[len(es)/2]
	ec := es[len(es)-1]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.RemoveEntity(ea)
		w.RemoveEntity(eb)
		w.RemoveEntity(ec)
		w.Flush()

		w.AddEntity(ea)
		w.AddEntity(eb)
		w.AddEntity(ec)
		w.Flush()
	}
}

// Todo benchmark with embedded world
