package main

import "fmt"

type Place struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type MapLayout struct {
	GraphNodes map[Place]struct{}
	Links      map[Place]map[Place]struct{}
}

func NewMapLayout() *MapLayout {
	return &MapLayout{
		GraphNodes: make(map[Place]struct{}),
		Links:      make(map[Place]map[Place]struct{}),
	}
}

func (mapLayout *MapLayout) AddPlace(place Place) bool {
	if _, exsits := mapLayout.GraphNodes[place]; exsits {
		return true
	}

	mapLayout.GraphNodes[place] = struct{}{}
	return true
}

func (mapLayout *MapLayout) AddLink(place1, place2 Place) {
	if _, exists := mapLayout.GraphNodes[place1]; !exists {
		mapLayout.AddPlace(place1)
	}

	if _, exists := mapLayout.GraphNodes[place2]; !exists {
		mapLayout.AddPlace(place2)
	}

	if _, exists := mapLayout.Links[place1]; !exists {
		mapLayout.Links[place1] = make(map[Place]struct{})
	}
	mapLayout.Links[place1][place2] = struct{}{}

}

func (mapLayout *MapLayout) PrintLinks() {
	root := Place{"Algeria", 3, 28}
	fmt.Printf("Printing all links adjacent to %s\n", root.Name)

	for node := range mapLayout.Links[root] {
		fmt.Printf("Link: %s -> %s\n", root.Name, node.Name)
	}

	fmt.Println("Printing all links.")
	for root, m := range mapLayout.Links {
		for vertex := range m {
			fmt.Printf("Link: %s -> %s\n", root.Name, vertex.Name)
		}
	}
}

func main() {
	mapLayout := NewMapLayout()
	root := Place{"Algeria", 3, 28}
	netherlands := Place{"Netherlands", 5.75, 52.5}
	korea := Place{"Korea", 124.1, -8.36}
	tunisia := Place{"Tunisia", 9, 34}

	mapLayout.AddPlace(root)
	mapLayout.AddPlace(netherlands)
	mapLayout.AddPlace(korea)
	mapLayout.AddPlace(tunisia)

	mapLayout.AddLink(root, netherlands)
	mapLayout.AddLink(root, korea)
	mapLayout.AddLink(root, tunisia)

	singapore := Place{"Singapore", 103.8, 1.36}
	uae := Place{"UAE", 54, 24}
	japan := Place{"Japan", 139.75, 35.68}

	mapLayout.AddLink(korea, singapore)
	mapLayout.AddLink(korea, japan)
	mapLayout.AddLink(netherlands, uae)
	mapLayout.PrintLinks()
}
