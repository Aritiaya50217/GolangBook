package main

import "fmt"

func swap(array []int, a int, b int) {
	array[a] = array[a] + array[b]
	array[b] = array[a] - array[b]
	array[a] = array[a] - array[b]
}

func shellSort(array []int, length int) {
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			var j = i
			for {
				if j-gap < 0 || array[j] >= array[j-gap] {
					break
				}
				swap(array, j, j-gap)
				j = j - gap
			}
		}
	}
}

func main() {
	scores := []int{9, 6, 5, 8, 0, 7, 4, 3, 1, 2}
	length := len(scores)
	shellSort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

/*
	First sorting
	gap = array.length / 2 = 5 (ระยะห่าง)
		9 > 7 swap
		i=0				   i=5
	[	9	6	5	8	0	7	4	3	1	2	]

			6 > 4 swap
			i=1				   i=6
	[	9	6	5	8	0	7	4	3	1	2	]

				5 > 3 swap
				i=2				   i=7
	[	9	6	5	8	0	7	4	3	1	2	]

					8 > 1 swap
					i=3				   i=8
	[	9	6	5	8	0	7	4	3	1	2	]

						0 < 2 no swap
						i=4				   i=9
	[	9	6	5	8	0	7	4	3	1	2	]

	after first sort
	[	7	4	3	1	0	9	6	5	8	2	]

	Second sorting
	gap = 5 / 2 = 2 (ระยะห่างรอบแรก / 2 = ระยะห่างรอบสอง)
	
			swap and sorting 0 3 6 7 8
		|		|		|		|		|
	[	7	4	3	1	0	9	6	5	8	2	]
			|		|		|		|		|
				swap and sorting 1 2 4 5 9 

	After
	[	0	1	3	2	6	4	7	5	8	9	]
			
	Third sorting
	gap = 2 / 2 = 1 (sort 1 ระยะ)

	[	0	1	3	2	6	4	7	5	8	9	]

	After
	[	0	1	2	3	4	5	6	7	8	9	]
			

*/