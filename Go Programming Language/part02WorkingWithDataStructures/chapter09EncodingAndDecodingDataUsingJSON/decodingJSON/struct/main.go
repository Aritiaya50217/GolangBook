package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Firstname string
	Lastname  string
}

func main() {
	var person People
	jsonStr := `{"firstname":"Wei-Meng","lastname":"Lee"}`
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err == nil {
		fmt.Println("Firstname : ",person.Firstname)
	} else {
		fmt.Println(err)
	}
}
