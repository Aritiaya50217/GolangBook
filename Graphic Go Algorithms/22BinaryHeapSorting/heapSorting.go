package main

import "fmt"

/*  Heap Sorting
	1. Build a heap
	2. After outputting the top element of the heap , adjust from top to bottom
	compare the top element with the root node of its left and right subtree and
	swap the smallest element to the top of the heap then adjust continuously
	until the leaf nodes to get new heap.

	- {10, 90, 20, 80, 30, 70, 40, 60, 50} build heap and then heap sort output.

	1. Initialize the heap and build the heap
						    10
						/		\
					 90			  20
				   /	\		 /	 \
				80		30	   70     40
			  /   \
			60     50

	Not LeafNode = 80 > left = 60	, 80 > right = 50 No need to move

 							10
						/		\
					 90			  20
				   /	\		 /	 \
				80		30	   70     40
			  /   \
			60     50

	Not LeafNode = 20 < left = 70	, 70 > right = 40 , 20 swap with 70

 							10
						/		\
					 90			  70
				   /	\		 /	 \
				80		30	   20     40
			  /   \
			60     50

	Not LeafNode = 90 > left = 80, 80 > right = 30 No need to move

 							10
						/		\
					 90			  70
				   /	\		 /	 \
				80		30	   20     40
			  /   \
			60     50

	Not LeafNode = 90 > left = 80, 80 > right = 30 No need to move

 							10
						/		\
					 90			  70
				   /	\		 /	 \
				80		30	   20     40
			  /   \
			60     50

	Not LeafNode = 10 < left = 90, 90 > right = 70 , 10 swap with 90

 							90
						/		\
					 10			  70
				   /	\		 /	 \
				80		30	   20     40
			  /   \
			60     50

	Still Not Leaf Node = 10 < left=80,80 > night=30,10 swap with 80

 							90
						/		\
					 80			  70
				   /	\		 /	 \
				10		30	   20     40
			  /   \
			60     50

	Still Not Leaf Node = 10 < left=80,80 > night=30,10 swap with 80

 							90
						/		\
					 80			  70
				   /	\		 /	 \
				10		30	   20     40
			  /   \
			60     50

	Still Not Leaf Node = 10 < left=60,60 > night=50,10 swap with 60

 							90
						/		\
					 80			  70
				   /	\		 /	 \
				60		30	   20     40
			  /   \
			10     50

	Create the heap finished

	2.Start heap sorting

							90
						/		\
					 80			  70
				   /	\		 /	 \
				60		30	   20     40
			  /   \
			10     50

	root = 90 and tail = 50 are exchanged
						   50
						/		\
					 80			  70
				   /	\		 /	 \
				60		30	   20     40
			  /
			10     90

	adjust the haep
	 						80
						/		\
					 60			  70
				   /	\		 /	 \
				50		30	   20     40
			  /
			10     90

	root = 80 and tail = 10 are exchanged
							10
						/		\
					 60			  70
				   /	\		 /	 \
				50		30	   20     40

			80     90

	adjust the haep
						   70
						/		\
					 60			  10
				   /	\		 /	 \
				50		30	   20     40

			80     90

	  						70
						/		\
					 60			  40
				   /	\		 /	 \
				50		30	   20     10

			80     90

	root = 70 and tail = 10 are exchanged

	  						10
						/		\
					 60			  40
				   /	\		 /
				50		30	   20     70

			80     90

	adjust the heap
	  						10
						/		\
					 60			  40
				   /	\		 /
				50		30	   20     70

			80     90

	  						60
						/		\
					 10			  40
				   /	\		 /
				50		30	   20     70

			80     90

							60
						/		\
					 50			  40
				   /	\		 /
				10		30	   20     70

			80     90

	root = 60 and tail = 20 are exchanged
							20
						/		\
					 50			  40
				   /	\
				10		30	   60     70

			80     90

	adjust the heap
							50
						/		\
					 20			  40
				   /	\
				10		30	   60     70

			80     90

							50
						/		\
					 30			  40
				   /	\
				10		20	   60     70

			80     90

	root = 50 and tail = 20 are exchanged

							20
						/		\
					 30			  40
				   /
				10		50	   60     70

			80     90

	adjust the heap
							40
						/		\
					 30			  20
				   /
				10		50	   60     70

			80     90

	root = 40 and tail = 10 are exchanged

							10
						/		\
					 30			  20

				40		50	   60     70

			80     90

	adjust the heap
							30
						/		\
					 10			  20

				40		50	   60     70

			80     90

	root = 30 and tail = 20 are exchanged

							20
						/
					 10			  30

				40		50	   60     70

			80     90

	no need adjust the heap
							10

					 20			  30

				40		50	   60     70

			80     90

	heap sort result
	[ 10 20 30 40 50 60 70 80 90 ]

*/

// Adjustment heap
func adjustHeap(array []int, currentIndex int, maxLength int) {
	// current non-leaf node
	var noLeafValue = array[currentIndex]
	// 2 * currentIndex+1 Current left subtree subscript
	for j := 2*currentIndex + 1; j <= maxLength; j = currentIndex*2 + 1 {
		if j < maxLength && array[j] < array[j+1] {
			// j Large subscript
			j++
		}
		if noLeafValue >= array[j] {
			break
		}
		// move up to the parent node
		array[currentIndex] = array[j]
		currentIndex = j
	}
	// to put in the position
	array[currentIndex] = noLeafValue
}

// Initialize the heap
func createHeap(array []int, length int) {
	// build a heap ,(length-1) / 2 scan half of the nodes with child nodes
	for i := (length - 1) / 2; i >= 0; i-- {
		adjustHeap(array, i, length-1)
	}
}

func heapSort(array []int, length int) {
	for i := length - 1; i > 0; i-- {
		var temp = array[0]
		array[0] = array[i]
		array[i] = temp
		adjustHeap(array, 0, i-1)
	}
}

func main() {
	var scores = []int{10, 90, 20, 80, 30, 70, 40, 60, 50}
	var length = len(scores)

	fmt.Printf("Before building a heap : \n")
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("After building a heap : \n")
	createHeap(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
	fmt.Printf("\n\n")

	fmt.Printf("After heap sorting : \n")
	heapSort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}
