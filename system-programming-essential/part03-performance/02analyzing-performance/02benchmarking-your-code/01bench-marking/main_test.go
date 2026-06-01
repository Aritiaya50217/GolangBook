package main

import "testing"

// Benchmark for the Sum function
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

// Sub-benchmark for the Sum function with different cases
func BenchmarkSumSub(b *testing.B) {
	cases := []struct {
		name string
		a, b int
	}{
		{"small", 1, 2},
		{"medium", 250, 500},
		{"large", 1000, 2000},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(c.a, c.b)
			}
		})
	}
}
