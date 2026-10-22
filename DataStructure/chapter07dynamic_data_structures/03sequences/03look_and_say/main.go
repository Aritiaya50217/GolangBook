package main

import (
	"fmt"
	"strconv"
)

func lock_say(str string) (rstr string) {
	cbyte := str[0]
	inc := 1

	for i := 1; i < len(str); i++ {
		dbyte := str[i]
		if dbyte == cbyte {
			inc++
			continue
		}
		rstr = rstr + strconv.Itoa(inc) + string(cbyte)
		cbyte = dbyte
		inc = 1
	}

	return rstr + strconv.Itoa(inc) + string(cbyte)
}

func main() {
	str := "1"
	fmt.Println("str : ", str)

	for i := 0; i < 8; i++ {
		str = lock_say(str)
		fmt.Println(str)
	}
}
