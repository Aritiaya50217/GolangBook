package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Structs can also be embedded within other structs.
// This replaces inheritance by simply storing the
// data type as another variable.
type Hacker struct {
	Person           Person
	FavoriteLanguage string
}

func main() {
	data := &Person{Name: "NanoDano", Age: 99}
	fmt.Println(data)

	hacker := &Hacker{
		Person:           *data,
		FavoriteLanguage: "GO",
	}

	fmt.Println(hacker)
	fmt.Println(hacker.Person.Name)
	fmt.Println(hacker)
}
