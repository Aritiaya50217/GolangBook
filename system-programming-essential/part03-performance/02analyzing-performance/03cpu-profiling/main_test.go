package main

import "testing"

func BenchmarkHeavyWork(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heavyWork()
	}
}
