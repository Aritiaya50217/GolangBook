package main

import "fmt"

type Name string

type SocialGraph struct {
	GraphNodes map[Name]struct{}
	Links      map[Name]map[Name]struct{}
}

func NewSocialGraph() *SocialGraph {
	return &SocialGraph{
		GraphNodes: make(map[Name]struct{}),
		Links:      make(map[Name]map[Name]struct{}),
	}
}

func (socialGraph *SocialGraph) AddEntity(name Name) bool {
	if _, exists := socialGraph.GraphNodes[name]; exists {
		return true
	}

	socialGraph.GraphNodes[name] = struct{}{}

	return true
}

func (socialGraph *SocialGraph) AddLink(name1, name2 Name) {
	if _, exists := socialGraph.GraphNodes[name1]; !exists {
		socialGraph.AddEntity(name1)
	}

	if _, exists := socialGraph.GraphNodes[name2]; !exists {
		socialGraph.AddEntity(name2)
	}

	if _, exists := socialGraph.Links[name1]; !exists {
		socialGraph.Links[name1] = make(map[Name]struct{})
	}

	socialGraph.Links[name1][name2] = struct{}{}
}

func (socialGraph *SocialGraph) PrintLinks() {
	root := Name("Root")
	fmt.Printf("Printing all links adjacent to %v\n", root)

	for node := range socialGraph.Links[root] {
		// Edge exists from u to v.
		fmt.Printf("Link: %v -> %v\n", root, node)

		fmt.Println("Printing all links.")

		for root, m := range socialGraph.Links {
			for vertex := range m {
				fmt.Printf("Link: %v -> %v\n", root, vertex)
			}
		}
	}
}

func main() {
	socialGraph := NewSocialGraph()
	root := Name("Root")
	john := Name("John Smith")
	per := Name("Per Jambeck")
	cynthia := Name("Cynthia Gibas")

	socialGraph.AddEntity(root)
	socialGraph.AddEntity(john)
	socialGraph.AddEntity(per)
	socialGraph.AddEntity(cynthia)

	socialGraph.AddLink(root, john)
	socialGraph.AddLink(root, per)
	socialGraph.AddLink(root, cynthia)

	mayo := Name("Mayo Smith")
	lorrie := Name("Lorrie Jambeck")
	ellie := Name("Ellie Vlocksen")

	socialGraph.AddLink(john, mayo)
	socialGraph.AddLink(john, lorrie)
	socialGraph.AddLink(per, ellie)
	socialGraph.PrintLinks()
}
