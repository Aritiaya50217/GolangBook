package main

import "fmt"

func main() {
	scores := []int{30, 40, 50, 70, 85, 90, 100}
	length := len(scores)
	searchValue := 40
	position := binarySearch(scores, length, searchValue)
	fmt.Printf("%d position : %d", searchValue, position)
	fmt.Printf("\n======================\n")

	searchValue = 90
	position = binarySearch(scores, length, searchValue)
	fmt.Printf("%d position : %d", searchValue, position)
}

func binarySearch(arrays []int, length int, searchValue int) int {
	low := 0   
	high := length
	mid := 0

	for {
		if low >= high {
			break
		}
		mid = (low + high) / 2
		if arrays[mid] == searchValue {
			return mid
		} else if arrays[mid] < searchValue {
			low = mid + 1
		} else if arrays[mid] > searchValue {
			high = mid - 1
		}
	}
	return -1
}

/*  searchValue = 40

	   low=0															high=6
	   							mid=(low+high)/2=3
								40 < array[mid] = 70
								high = mid = 3
	[	30		40		50				70				85		90		100	]
	
	   low=0						  		  		  high=3
				mid=(low+high)/2=1
				40 == array[mid]
				Found 40 index = 1
	[	30				40				50				70				85		90		100	]


	searchValue = 90

		low=0																			high=6
												mid=(low+high)/2=3
												90 > array[mid]=70
												low=mid=3
	[	30				40				50				70				85		90		100	]

													  low=3										high=6
													  				mid=(low_high)/2 = 4
																	90 > array[mid]=85
																	low=mid=4
	[	30				40				50				70				85				90		100	]

																		low=4							high=6
																				mid=(low+high)/2=5
																				90 == array[mid]
																				Found 90 index = 5

	[	30				40				50				70				85				90				100	]


*/

