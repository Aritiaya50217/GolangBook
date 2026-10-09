package main

import "fmt"

func main() {
	var k, l, m int
	var arr [5][5][5]int
	// การเปรียบเทียบทุกทรีเปิลของสมาชิกในอาร์เรย์
	for k = 0; k < 5; k++ {
		for l = 0; l < 5; l++ {
			for m = 0; m < 5; m++ {
				arr[k][l][m] = 1
				fmt.Println("Element Value ", k, l, m, " is", arr[k][l][m])
			}
		}
	}
}
