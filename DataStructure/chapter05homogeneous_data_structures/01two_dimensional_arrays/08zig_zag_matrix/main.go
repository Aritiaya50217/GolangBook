package main

import "fmt"

func PrintZigZag(n int) []int {
	zigzag := make([]int, n*n)
	i := 0
	m := n * 2
	for p := 1; p <= m; p++ {
		x := p - n
		if x < 0 {
			x = 0
		}

		y := p - 1
		if y > n - 1 {
			y = n - 1
		}

		j := m - p
		if j > p {
			j = p
		}

		for k := 0; k < j; k++ {
			if p&1 == 0 {
				zigzag[(x+k)*n+y-k] = i
			} else {
				zigzag[(y-k)*n+x+k] = i
			}
			i++
		}
	}
	return zigzag
}

func main() {
	n := 5
	length := 2
	for i, sketch := range PrintZigZag(n) {
		fmt.Printf("%*d ", length, sketch)
		if i%n == n-1 {
			fmt.Println("")
		}
	}
}
