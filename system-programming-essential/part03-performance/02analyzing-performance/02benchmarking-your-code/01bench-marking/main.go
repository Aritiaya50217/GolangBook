package main

func Fib(n int) int {
	if n <= 1 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func Sum(a, b int) int {
	return a + b
}
