package thedifferencebetweenatestandabenchmark

import "testing"

type Engine struct{}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

func BenchmarkAdd(b *testing.B) {
	e := Engine{}

	// run the Add function b.N times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}
}
