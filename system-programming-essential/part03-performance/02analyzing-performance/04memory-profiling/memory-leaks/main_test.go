package main

import "testing"

func BenchmarkMemoryLeak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = make([]byte, 1024)
	}
}
