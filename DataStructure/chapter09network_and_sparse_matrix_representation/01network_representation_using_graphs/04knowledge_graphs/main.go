package main

import "fmt"

type Class struct {
	Name string
}

type KnowledgeGraph struct {
	GraphNodes map[Class]struct{}
	Links      map[Class]map[Class]struct{}
}

func NewKnowledgeGraph() *KnowledgeGraph {
	return &KnowledgeGraph{
		GraphNodes: make(map[Class]struct{}),
		Links:      make(map[Class]map[Class]struct{}),
	}
}

func (knowledgeGraph *KnowledgeGraph) AddClass(class Class) bool {
	if _, exists := knowledgeGraph.GraphNodes[class]; exists {
		return true
	}

	knowledgeGraph.GraphNodes[class] = struct{}{}
	return true
}

func (knowledgeGraph *KnowledgeGraph) AddLink(class1, class2 Class) {
	if _, exists := knowledgeGraph.GraphNodes[class1]; !exists {
		knowledgeGraph.AddClass(class1)
	}

	if _, exists := knowledgeGraph.GraphNodes[class2]; !exists {
		knowledgeGraph.AddClass(class2)
	}

	if _, exists := knowledgeGraph.Links[class1]; !exists {
		knowledgeGraph.Links[class1] = make(map[Class]struct{})
	}
	knowledgeGraph.Links[class1][class2] = struct{}{}
}

func (knowledgeGraph *KnowledgeGraph) PrintLinks() {
	car := Class{"Car"}

	fmt.Printf("Printing all links adjacent to %s\n", car.Name)

	for node := range knowledgeGraph.Links[car] {
		fmt.Printf("Link: %s -> %s\n", car.Name, node.Name)
	}

	fmt.Println("Printing all links.")

	for car, m := range knowledgeGraph.Links {
		for vertex := range m {
			fmt.Printf("Link: %s -> %s\n", car.Name, vertex.Name)
		}
	}
}

func main() {
	knowledgeGraph := NewKnowledgeGraph()
	car := Class{"Car"}
	tyre := Class{"Tyre"}
	door := Class{"Door"}
	hood := Class{"Hood"}

	knowledgeGraph.AddClass(car)
	knowledgeGraph.AddClass(tyre)
	knowledgeGraph.AddClass(door)
	knowledgeGraph.AddClass(hood)

	knowledgeGraph.AddLink(car, tyre)
	knowledgeGraph.AddLink(car, door)
	knowledgeGraph.AddLink(car, hood)

	tube := Class{"Tube"}
	axle := Class{"Axle"}
	handle := Class{"Handle"}
	windowGlass := Class{"Window Glass"}

	knowledgeGraph.AddClass(tube)
	knowledgeGraph.AddClass(axle)
	knowledgeGraph.AddClass(handle)
	knowledgeGraph.AddClass(windowGlass)

	knowledgeGraph.AddLink(tyre, tube)
	knowledgeGraph.AddLink(tyre, axle)
	knowledgeGraph.AddLink(door, handle)
	knowledgeGraph.AddLink(door, windowGlass)
	knowledgeGraph.PrintLinks()
}
