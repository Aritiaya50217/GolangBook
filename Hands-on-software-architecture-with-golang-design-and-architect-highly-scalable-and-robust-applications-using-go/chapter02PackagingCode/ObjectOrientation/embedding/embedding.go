package main

import "fmt"

type Bird struct {
	featherLength  int
	classification string
}

type Pigeon struct {
	Bird
	Name string
}

func main() {
	p := Pigeon{
		Name: "Tweety",
	}
	p.featherLength = 10
	fmt.Println(p)

}
