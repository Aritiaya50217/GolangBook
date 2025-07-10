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
	var persons []People
	jsonStr := `[
					{
						"firstname":"Wei-Meng",
						"lastname":"Lee"
					},
					{
						"firstname":"Mickey",
						"lastname":"Mouse"
					}
				]`
	json.Unmarshal([]byte(jsonStr), &persons)

	for _, person := range persons {
		fmt.Printf("Firstname : %s Lastname : %s\n", person.Firstname, person.Lastname)
	}

}
