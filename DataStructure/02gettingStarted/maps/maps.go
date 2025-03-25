package main

import "fmt"

func main() {
	var languages = map[int]string{
		3: "English",
		4: "French",
		5: "Spanish",
	}

	products := make(map[int]string)
	products[1] = "chair"
	products[2] = "table"

	for i, value := range languages {
		fmt.Printf("language %d : %s\n", i, value)
	}
	fmt.Println("products : ", products)

	fmt.Println("product with key 2", products[2])
	fmt.Println(products[2])

	delete(products, 1) // delete "chair"
	fmt.Println("After : ", products)
}
