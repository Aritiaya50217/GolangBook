package main

import "fmt"

// findElementsWithSum of k from arr of size
func findElementsWithSum(arr [10]int, combinations [19]int, size, k, addValue, l, m int) int {
	var num int = 0
	if addValue > k {
		return -1
	}

	if addValue == k {
		num += 1
		for p := 0; p < m; p++ {
			fmt.Printf("%d", arr[combinations[p]])
		}
		fmt.Println(" ")
	}

	for i := l; i < size; i++ {
		combinations[m] = 1
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l += 1
	}
	return num
}

// example การจัดเรียง N ราชินีในกระดานหมากรุก NxN โดยไม่ให้ราชินีสองตัวอยู่ในตำแหน่งที่สามารถโจมตีกันได้
func isSafe(board [][]int, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	// ตรวจสอบแนวทแยง
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}

func solveNQueens(board [][]int, row, n int) bool {
	if row == n {
		return true // เจอคำตอบ
	}

	for col := 0; col < n; col++ {
		if isSafe(board, row, col, n) {
			board[row][col] = 1 // วาง queen ตำแหน่งนี้

			if solveNQueens(board, row+1, n) {
				return true
			}

			board[row][col] = 0 // ถ้าไม่สำเร็จย้อนกลับมา
		}
	}
	return false // กรณีไม่สำเร็จ
}

func printBoard(board [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 1 {
				fmt.Print("Q ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {

	arr := [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	addedSum := 18
	combinations := [19]int{}

	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)

	// example
	n := 4
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	if solveNQueens(board, 0, n) {
		printBoard(board, n)
	} else {
		fmt.Println("Solution does not exist")
	}

}
