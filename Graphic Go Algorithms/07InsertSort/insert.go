package main

import "fmt"

func main() {
	// index start from 0
	scores := []int{80, 70, 60, 50, 95}
	length := len(scores)
	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}

}

func sort(arrays []int, length int) {
	for i := 0; i < length; i++ {
		// take unsorted new elements
		insertElement := arrays[i]
		// inserted position
		insertPosition := i
		for j := insertPosition - 1; j >= 0; j-- {
			/* if the new element is smaller than the sorted element ,
			it is shifted to the right
			*/
			if insertElement < arrays[j] {
				arrays[j+1] = arrays[j]
				insertPosition--
			}
		}
		// insert the new element
		arrays[insertPosition] = insertElement
	}
}

/*
	First sorting

			insertElement ( 80 > 70 = ? ถ้ามากกว่า swap )
	[	80		70		60		50		95]

	[	70		80		60		50		95]

	Second sorting
					insertElement 
	[	70		80		60		50		95]

	( 70 > 60 and 80 > 60 )
	[	60		70		80		50		95]

	Third sorting
							insertElement 
	[	70		80		60		50		95]

	( 60 > 50 , 70 > 50 , 80 > 50 )
	[	50		60		70		80		95]

	Forth sorting
									insertElement 
	[	70		80		60		50		95]

	( 50 < 95,60 < 95 , 70 < 95 , 80 < 95 )
	[	50		60		70		80		95]

*/
