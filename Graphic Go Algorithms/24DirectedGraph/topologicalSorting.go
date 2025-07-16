package main

import "fmt"

const MAX_VERTEX_SIZE = 5
const STACKSIZE = 1000

type Vertex struct {
	data    string
	visited bool // have you visited
}

// Stack save current vertices
var top = -1
var stacks [STACKSIZE]int

func push(element int) {
	top++
	stacks[top] = element
}

func pop() int {
	if top == -1 {
		return -1
	}
	var data = stacks[top]
	top--
	return data
}

func peek() int {
	if top == -1 {
		return -1
	}
	var data = stacks[top]
	return data
}

func isEmpty() bool {
	if top <= -1 {
		return true
	}
	return false
}

var size = 0 // current  vertex size
var vertexs [MAX_VERTEX_SIZE]Vertex

// An array of topological sort results , recording the sequence number of each node
var topologys [MAX_VERTEX_SIZE]Vertex
var adjacencyMatrix [MAX_VERTEX_SIZE][MAX_VERTEX_SIZE]int

func addVertex(data string) {
	var vertex Vertex
	vertex.data = data
	vertex.visited = false
	vertexs[size] = vertex
	size++
}

// add adjacent edges
func addEdge(from int, to int) {
	// A -> B = B -> A
	adjacencyMatrix[from][to] = 1
}

func removeVertex(vertex int) {
	if vertex != size-1 {
		// if the vertex is the last element ,  the end
		for i := vertex; i < size-1; i++ { // The vertices are removed from the vertex array
			vertexs[i] = vertexs[i+1]
		}
		for row := vertex; row < size-1; row++ {
			// move up a row
			for col := 0; col < size-1; col++ {
				adjacencyMatrix[row][col] = adjacencyMatrix[row+1][col]
			}
		}
		for col := vertex; col < size-1; col++ { // move left a row
			for row := 0; row < size-1; row++ {
				adjacencyMatrix[row][col] = adjacencyMatrix[row][col+1]
			}
		}
	}
	// Decrease the number of vertices
	size--
}

func topologySort() {
	for {
		if size <= 0 {
			break
		}
		var noSuccessorVertex = getNoSuccessorVertex()
		if noSuccessorVertex == -1 {
			fmt.Printf("There is ring in Graph \n")
			return
		}
		// copy the deleted node to the sorted array
		topologys[size-1] = vertexs[noSuccessorVertex]
		// Delete no successor node
		removeVertex(noSuccessorVertex)
	}
}

func getNoSuccessorVertex() int {
	var existSuccessor = false
	for row := 0; row < size; row++ {
		// for each vertex
		existSuccessor = false
		/** If the node has a fixed row, each column has a 1, indicating that
		 the node has a successor, terminating the loop
		**/
		for col := 0; col < size; col++ {
			if adjacencyMatrix[row][col] == 1 {
				existSuccessor = true
				break
			}
		}
		if !existSuccessor {
			// If the node has no successor , return it's subscript
			return row
		}
	}
	return -1
}

func printGraph() {
	fmt.Printf("Two-dimensional array traversal vertex edge and adjacent array : \n ")
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertexs[i].data)
	}
	fmt.Printf("\n")
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s ", vertexs[i].data)
		for j := 0; j < MAX_VERTEX_SIZE; j++ {
			fmt.Printf("%d ", adjacencyMatrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {
	addVertex("A")
	addVertex("B")
	addVertex("C")
	addVertex("D")
	addVertex("E")

	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(0, 3)
	addEdge(1, 2)
	addEdge(1, 3)
	addEdge(2, 3)
	addEdge(3, 4)

	// Two-dimensional array traversal output vertex edge and adjacent array
	printGraph()
	fmt.Printf("\nDepth-First Search traversal output : \n")
	fmt.Printf("Directed Graph Topological Sorting : \n")
	topologySort()
	for i := 0; i < MAX_VERTEX_SIZE; i++ {
		fmt.Printf("%s-> ", topologys[i].data)
	}
}

/*	1. The adjacency matrix is described above:
	The total number of vertices is a two-dimensional array size, if have value of
	the edge is 1, otherwise no value of the edge is 0.

			A	B	C	D	E
		A	0 	1	1	1	0
		B	0	0	1	1	0
		C	0	0	0	1	0
		D	0	0	0	0	1
		E	0	0	0	0	0

	Topological sorting	from vertex A : A -> B -> C -> D -> E

	find no successor vertices E then save to topologys, last E remove from the graph

		vetex												topologys
		A		---->		B									E
		|		 			|
edge	|					|
		v			  		v
		C		---->		D

	find on successor vertices D then save to topologys, last D remove from the graph

		vetex									topologys
		A	--->  B									D
		|	  	/									E
edge	|	   /
		v	  /
		C

	find on successor vertices C then save to topologys , last C remove from the graph

		vetex									topologys
		A	--->  B									C
													D
				  									E

	find no successor vertices B then save to topologys , last B remove from the graph

		vetex									topologys
		A											B
													C
													D
				  									E

	find no successor vertices A then save to topologys , last A remove from the graph

		topologys
			A
			B
			C
			D
			E
	output
	A -> B -> C -> D -> E


*/
