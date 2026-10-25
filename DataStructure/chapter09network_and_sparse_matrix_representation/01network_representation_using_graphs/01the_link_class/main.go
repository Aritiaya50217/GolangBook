package main

import "fmt"

type SocialGraph struct {
	Size  int
	Links [][]Link
}

type Link struct {
	Vertex1    int
	Vertex2    int
	LinkWeight int
}

func NewSocialGraph(num int) *SocialGraph {
	return &SocialGraph{
		Size:  num,
		Links: make([][]Link, num),
	}
}

func (socialGraph *SocialGraph) AddLink(vertex1, vertex2, weight int) {
	socialGraph.Links[vertex1] = append(socialGraph.Links[vertex1],
		Link{
			Vertex1:    vertex1,
			Vertex2:    vertex2,
			LinkWeight: weight,
		})
}

func (socialGraph *SocialGraph) PrintLinks() {
	vertex := 0

	fmt.Printf("Printing all links from %d\n", vertex)
	
	for _, link := range socialGraph.Links[vertex] {
		fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
	}

	fmt.Println("Printing all links in graph.")

	for _, adjacent := range socialGraph.Links {
		for _, link := range adjacent {
			fmt.Printf("Link: %d -> %d (%d)\n", link.Vertex1, link.Vertex2, link.LinkWeight)
		}
	}
}

func main() {
	var socialGraph *SocialGraph
	socialGraph = NewSocialGraph(4)

	socialGraph.AddLink(0, 1, 1)
	socialGraph.AddLink(0, 2, 1)
	socialGraph.AddLink(1, 3, 1)
	socialGraph.AddLink(2, 4, 1)
	socialGraph.PrintLinks()

}
