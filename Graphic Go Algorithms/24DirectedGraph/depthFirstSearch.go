package main

import "fmt"

const MAX_VERTEX_SIZE = 5
const STACKSIZE = 1000

type Vertex struct {
	data    string
	visited bool // have you visited
}

var top = -1 // stack saves current vertices
var stacks = make([]int, STACKSIZE)

func push(element int) {
	// element ใหม่ อยู่ด้านบน
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

var size = 0 // current vertex size
var vertexs = make([]Vertex, MAX_VERTEX_SIZE)
// adjacent = ชิด , ใกล้กัน
var adjacencyMatrix [MAX_VERTEX_SIZE][MAX_VERTEX_SIZE]int

func addVertex(data string) {
	var vertex Vertex
	vertex.data = data
	vertex.visited = false
	vertexs[size] = vertex
	size++
}

// Add adjacent edges
func addEdge(from int, to int) {
	// A -> B != B -> A
	// if have value of the edge is 1
	adjacencyMatrix[from][to] = 1
}

func clear() {
	for i := 0; i < size; i++ {
		vertexs[i].visited = false
	}
}

func depthFirstSearch() {
	// Start searching from the first vertix
	vertexs[0].visited = true
	fmt.Printf("%s ", vertexs[0].data)
	push(0)
	for {
		if isEmpty() {
			break
		}
		var row = peek()
		// get adjacent vertex position that have not been visited
		var col = findAdjacencyUnVisitedVertex(row)
		if col == -1 {
			pop()
		} else {
			vertexs[col].visited = true
			fmt.Printf("-> %s ", vertexs[col].data)
			push(col)
		}
	}
	clear()
}

// get adjacent vertex positions that have not been visited
func findAdjacencyUnVisitedVertex(row int) int {
	for col := 0; col < size; col++ {
		if adjacencyMatrix[row][col] == 1 && !vertexs[col].visited {
			return col
		}
	}
	return -1
}

func printGraph() {
	fmt.Printf("Two-dimensional array traversal vertex edge and adjecent array : \n  ")
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

	fmt.Printf("\nDepth-first search traversal output : \n")
	depthFirstSearch()

}

/* Directed Grap has direction : A -> B and B -> A are different

	vetex
		A		---->		B
		|		 			|
edge	|					|	  E
		v			  		v	/
		C		---->		D


- The adjacency matrix is described above
The total number of vertices is a two-dimentional array size , if have value of the edge is 1
otherwise no value of the edge is 0

		A	B	C	D	E										  Stack
	A	0 -	1	1	1	0		A -> B -> C -> D				[	E	]
			  \
	B	0	0	1	1	0		B -> C -> D						[	D	]
				  \
	C	0	0	0	1	0		C -> D							[	C	]
					  \
	D	0	0	0	0	1		D -> E							[	B	]

	E	0	0	0	0	0		E								[	A	]


*/
