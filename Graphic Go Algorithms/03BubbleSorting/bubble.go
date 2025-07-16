package main

import "fmt"

func main() {
	// index starts from 0
	scores := []int{90, 70, 50, 80, 60, 85}
	length := len(scores)
	fmt.Println("Before : ", scores)
	sort(scores, length)

	fmt.Println("After : ")

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func sort(arrays []int, length int) {
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			// arrays[j] => ตัวเลขที่ใช้เทียบ
			// arrays[j+1] => ตัวเลขถัดไป
			if arrays[j] > arrays[j+1] {
				flag := arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = flag
			}
		}
	}
}

/* First sorting
	60 > 50 (ถ้ามากกว่า swap)
[	60		50		95		80		70	 ]
[	50		60		...		95	 ]

Second sorting
	50 < 60 (ถ้าน้อยกว่า swap)
[ 	50		60		80	 	70		95	]
			60 < 80 (ถ้าน้อยกว่า swap)
[	50		60		80	... ]
					80 > 70 (ถ้ามากกว่า swap)
[	50		60		80		70	... ]


Third sorting
	50 < 60 (ถ้าน้อยกว่า swap)
[ 	50		60		70	 	80		95	]
			60 < 70 (ถ้าน้อยกว่า swap)
[	50		60		70		80		95 ]
*/
