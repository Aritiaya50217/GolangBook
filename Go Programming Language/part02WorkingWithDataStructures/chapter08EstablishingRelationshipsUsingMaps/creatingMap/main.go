package main

import (
	"fmt"
	"sort"
)

type keyValue struct {
	key   string
	value int
}

type kvPairs []keyValue

var heights map[string]int

func (p kvPairs) Len() int {
	// return the length of the collection
	return len(p)
}

func (p kvPairs) Less(i, j int) bool {
	return p[i].value < p[j].value
}

func (p kvPairs) Swap(i, j int) {
	// swap the items in the collection
	p[i], p[j] = p[j], p[i]
}

func main() {
	heights := make(map[string]int)
	heights["Peter"] = 170
	heights["Joan"] = 168
	heights["Jim"] = 150
	heights["Tommy"] = 110

	fmt.Println("heights : ", heights)

	// delete a key
	if _, ok := heights["Joan"]; ok {
		delete(heights, "Joan")
	} else {
		fmt.Println("Key does not exist")
	}
	fmt.Println("after deleted : ", heights)

	// get all the keys in a map
	var keys []string
	for key := range heights {
		keys = append(keys, key)
	}
	fmt.Println("keys : ", keys)

	// sort
	sort.Strings(keys)
	fmt.Println("keys : ", keys)

	p := make(kvPairs, len(heights))
	i := 0
	for k, v := range heights {
		p[i] = keyValue{k, v}
		i++
	}
	
	fmt.Println("p : ", p)
	fmt.Println(p.Len())
	fmt.Println(p.Less(1, 0))
	p.Swap(0, 1)
	fmt.Println("swap : ", p)

}
