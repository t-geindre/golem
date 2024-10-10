package test

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/pkg/golem"
	"github.com/t-geindre/golem/test/mock"
	"testing"
)

const (
	Size10K  = 1e4
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
			name: "Update complex world with 100K entities",
			args: args{world: mock.GetComplexWorld(Size100K)},
		},
		{
			name: "Update dead world with 100K entities",
			args: args{world: mock.GetDeadWorld(Size100K)},
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
			name: "Draw complex world with 100K entities",
			args: args{world: mock.GetComplexWorld(Size100K)},
		},
		{
			name: "Draw dead world with 100K entities",
			args: args{world: mock.GetDeadWorld(Size100K)},
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

func BenchmarkWorldRemoveEntities(b *testing.B) {
	type args struct {
		world          func() golem.World
		enCount        int
		immediateFlush bool
	}
	tests := []struct {
		name string
		args args
	}{
		// Use only simple world for this benchmark, all entities are on a single layer
		{
			name: "Remove 1 entity in simple 1M world, immediate flush",
			args: args{
				world:          func() golem.World { return mock.GetSimpleWorld(Size1M) },
				enCount:        1,
				immediateFlush: false,
			},
		},
		{
			name: "Remove 1 entity in simple 1M world, delayed flush",
			args: args{
				world:          func() golem.World { return mock.GetSimpleWorld(Size1M) },
				enCount:        1,
				immediateFlush: false,
			},
		},
		{
			name: "Remove 10K entity in simple 1M world, immediate flush",
			args: args{
				world:          func() golem.World { return mock.GetSimpleWorld(Size1M) },
				enCount:        Size10K,
				immediateFlush: false,
			},
		},
		{
			name: "Remove 10K entity in simple 1M world, delayed flush",
			args: args{
				world:          func() golem.World { return mock.GetSimpleWorld(Size1M) },
				enCount:        Size10K,
				immediateFlush: false,
			},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			w := tt.args.world()
			ens := w.GetEntities(mock.LayerA)
			if tt.args.enCount > len(ens) {
				b.Fatalf("Not enough entities in the world")
			}

			ens = ens[:tt.args.enCount]
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for _, en := range ens {
					w.RemoveEntity(en)
					if tt.args.immediateFlush {
						w.Flush()
					}
				}
			}

			if !tt.args.immediateFlush {
				w.Flush()
			}
		})
	}
}
