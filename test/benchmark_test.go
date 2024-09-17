package test

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/t-geindre/golem/test/mock"
	"testing"
)

func BenchmarkSimpleWorldUpdate(b *testing.B) {
	w := mock.GetSimpleWorld(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update()
	}
}

func BenchmarkSimpleWorldDraw(b *testing.B) {
	w := mock.GetSimpleWorld(100000)
	img := &ebiten.Image{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Draw(img)
	}
}

func BenchmarkComplexWorldUpdate(b *testing.B) {
	w := mock.GetComplexWorld(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update()
	}
}

func BenchmarkComplexWorldDraw(b *testing.B) {
	w := mock.GetComplexWorld(100000)
	img := &ebiten.Image{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Draw(img)
	}
}

func BenchmarkDeadWorldUpdate(b *testing.B) {
	w := mock.GetDeadWorld(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Update()
	}
}

func BenchmarkDeadWorldDraw(b *testing.B) {
	w := mock.GetDeadWorld(100000)
	img := &ebiten.Image{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Draw(img)
	}
}

func BenchmarkWorldAddRemoveSingleEntity(b *testing.B) {
	w := mock.GetDeadWorld(1000000)
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
	w := mock.GetDeadWorld(1000000)
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
