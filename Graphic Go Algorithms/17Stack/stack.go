package main

import "fmt"

// Stack  :  FILO (First In Last Out) sequence.
type Node struct {
	data string
	next *Node
}

var top *Node = nil
var size int

func push(element string) {
	if top == nil {
		top = new(Node)
		top.data = element
	} else {
		var newNode *Node = new(Node)
		newNode.data = element
		newNode.next = top
		top = newNode
	}
	size++
}

func pop() *Node {
	if top == nil {
		return nil
	}
	var p = top
	// top move down
	top = top.next
	p.next = nil
	size--
	return p
}

func output() {
	fmt.Printf("Top ")
	var node *Node = nil
	for {
		node = pop()
		if node == nil {
			break
		}
		fmt.Printf("%s -> ", node.data)
	}
	fmt.Printf("End\n")
}

func main() {
	push("A")
	push("B")
	push("C")
	push("D")

	output()
}

/*	Push A into Stack

		top=newNode("A")
			|
		   push
			|   [		]
			v	[		]
		top		[	A	]


	Push B into Stack

		newNode = newNode("B")
				B
		newNode.next = top	(เชื่อม node)
				|
			   push
			   	|	[		]
				v	[	B	]
					[	|	]
		top			[	A	]


	top = newNode 	[		]
		top			[	B	]
					[	|	]
					[	A	]


	Push C into Stack

		newNode = newNode("C")
				C
		newNode.next = top	(เชื่อม node)
				|
			   push
			   	|	[	C	]
				v	[	|	]
		top			[	B	]
					[	|	]
					[	A	]


	top = newNode
		top	   		[	C	]
					[	|	]
					[	B	]
					[	|	]
					[	A	]

	Pop C from Stack

					  p=top
		p -- top --	[	c	]
					[	|	]
					[	B	]
					[	|	]
					[	A	]


				 top = top.next
		p 			[	c	]
					[	|	]
		top 	  	[	B	]
					[	|	]
					[	A	]


				  p.next = nil
		p 				c	
					
		top 	  	[	B	]
					[	|	]
					[	A	]


*/
