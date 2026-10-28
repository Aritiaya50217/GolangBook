package main

import "fmt"

func addOne(num *int) {
	*num++
	fmt.Println("added to num", num, "Address of num", &num, "Value Points To", *num)
}

func main() {
	number := 17
	fmt.Println("value of number", number, "Address of number", &number)
	addOne(&number)
	fmt.Println("value of number after adding One", number, "Address of", &number)
}
