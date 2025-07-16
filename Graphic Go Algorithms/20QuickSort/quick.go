package main

import "fmt"

func main() {
	scores := []int{50, 65, 99, 87, 74, 63, 76, 100, 92}
	length := len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(array []int, length int) {
	if length > 0 {
		quickSort(array, 0, length-1)
	}
}

func quickSort(array []int, low int, high int) {
	if low > high {
		return
	}
	var i = low
	var j = high
	var threshold = array[low]
	// Alternately scanned from both ends of the list
	for {
		if i >= j {
			break
		}
		// Find the first position less than threshold from rigth to left
		for {
			if i >= j || array[j] <= threshold {
				break
			}
			j--
		}
		// replace the low with a smaller number than the threshold
		if i < j {
			array[i] = array[j]
			i++
		}
		// find the first position greater than threshold from left to right
		for {
			if i >= j || array[i] > threshold {
				break
			}
			i++
		}
		// Replace the high with a number larger than the threshold
		if i < j {
			array[j] = array[i]
			j--
		}
	}
	array[i] = threshold
	// left quickSort
	quickSort(array, low, i-1)
	// right quickSort
	quickSort(array, i+1, high)
}
