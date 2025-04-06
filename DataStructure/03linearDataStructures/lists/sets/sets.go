package main

import "fmt"

// set class
type Set struct {
	integerMap map[int]bool
}

// create the map of integer and bool
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// adds the element to the set
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	}
}

// deletes the element from the set
func (set *Set) DeleteElement(element int) {
	delete(set.integerMap, element)
}

// checks if element is in the set
func (set *Set) ContainsElement(element int) bool {
	exists := set.integerMap[element]
	return exists
}

func main() {
	set := Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println("sets : ", set)
	// delete
	set.DeleteElement(2)
	fmt.Println("After deleted : ", set)

	// check  element
	fmt.Println("2 in set : ", set.ContainsElement(2))

}
