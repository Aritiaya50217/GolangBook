package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	length := len(scores)
	sort(scores, length)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(arrays []int, length int) {
	var minIndex int
	for i := 0; i < length-1; i++ {
		minIndex = i
		// save the minimum value of each loop as the first element
		var minValue = arrays[minIndex]
		for j := i; j < length-1; j++ {
			if minValue > arrays[j+1] {
				minValue = arrays[j+1]
				minIndex = j + 1
			}
		}
		if i != minIndex {
			var temp = arrays[i]
			arrays[i] = arrays[minIndex]
			arrays[minIndex] = temp
		}
	}
}

/*  First sorting

	minIndex=0 , i = 0		j = 1 
	[	60				80			95			50			70	]

	minIndex=0 , i = 0			   j = 2
	[	60				80			95			50			70	]

	minIndex=0 , i = 0							j = 3  ( 50 < 60 ค่า minIndex จึงเลื่อนมาที่ 50)
	[	60				80			95			50			70	]

		i = 0									minIndex=3 	j = 4
	[	60				80			95			50			70	]

	i != minIndex and 60 > 50 swap (ตำแหน่ง i ไม่ใช่เลขที่น้อยที่สุด จึงต้องสลับกับเลขที่อยู่ลำดับแรก)
	minValue , i=0                             minIndex=3   j = 4
	[	50				80			95			60			70	]

	Second sorting
				minIndex=1 , i=1	j=2
	[	50				80			95			60			70	]

				minIndex=1 , i=1				j=3 
	[	50				80			95			60			70	]

						i=1					minIndex=3		j=4 (80 > 60 จึงเปลี่ยนมาใช้ 60 ในการเทียบ)
	[	50				80			95			60			70	]
	
	i != minIndex and 80 > 60 swap 
						i=1					minIndex=3     	j=4
	[	50				60			95			80			70	]
	
	Third sorting
							minIndex=2 , i=2	j=3
	[	50				60			95			80			70	]

									i=2		minIndex=3  	j=4
	[	50				60			95			80			70	]
	
									i=2						minIndex=4 ,j=4 (80 > 70 ค่า minIndex จึงเลื่อนมาที่ 70)
	[	50				60			95			80			70	]
	
									i != minIndex and 95 > 70 swap

									i=2						minIndex=4 ,j=4 
	[	50				60			70			80			95	]
	
	Forth sorting
										minIndex=3,i=3		j=4  
	[	50				60			70			80			95	]
	
										minIndex=3,i=3		j=4 	(80 < 95 ค่า minIndex จึงไม่มีการ swap)
	[	50				60			70			80			95	]

										i == minIndex no swap and finished
										minIndex=3,i=3		j=4 	
	[	50				60			70			80			95	]
	
*/