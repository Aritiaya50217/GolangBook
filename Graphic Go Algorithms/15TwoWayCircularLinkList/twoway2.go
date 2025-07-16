package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.prev = nil
	head.next = nil

	var nodeB *Node = &Node{data: "B", prev: head, next: nil}
	head.next = nodeB

	var nodeC *Node = &Node{data: "C", prev: nodeB, next: nil}
	nodeB.next = nodeC

	tail.data = "D"
	tail.prev = nodeC
	tail.next = head

	nodeC.next = tail
	head.prev = tail

}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.data = data
	// newNode next point to next node
	newNode.next = p.next
	// current next point to newNode
	p.next = newNode
	newNode.prev = p
	// newNode => next => prev
	newNode.next.prev = newNode
}

func output() {
	var p = head
	for {
		fmt.Printf("%s-> ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("End\n")

	p = tail
	for {
		fmt.Printf("%s ->", p.data)
		p = p.prev
		if p == tail {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("Start\n\n")
}

func main() {
	initial()
	fmt.Printf("Insert a new node E at index 2 : \n")
	insert(2, "E")
	output()
}

/*


			   head									p
								- next ->
				A	 								B
			|		|			- prev ->		|		|		  newNode
		  next	   prev						  prev	   next 		E
			|		|							|		|
								- prev ->
	tail		D									C
								- next ->



			   head									p
								- next ->
				A	 								B
			|		|			- prev ->		|		|	newNode.next = p.next
		  next	   prev						  prev	   next 	E
			|		|							|		|	   /
								- prev ->  					 next
															/
	tail		D									C
								- next ->



			   head									p
								- next ->	p.next = newNode
				A	 								B
														   \
															next
			|		|			- prev ->		|		|	  \
		  next	   prev						  prev	   next 	E
			|		|							|		|	   /
								- prev ->  					 next
															/
	tail		D									C
								- next ->


	   			head								p
								- next ->
				A	 								B
														   	\		\
								- prev ->					prev	next
			|		|							|		|	  	\		\
																newNode.prev = p
		  next	   prev						  prev	   next 		E
			|		|							|		|	  	  /
								- prev ->  					   next
															  /
	tail		D									C
								- next ->


			 head								p
								- next ->
				A	 								B
														   	\		\
								- prev ->					prev	next
			|		|							|		|	  	\		\
																newNode.next.prev = newNode
		  next	   prev						  prev	   next 		E
			|		|							|		|	  	  /		/
								- prev ->  					   prev	  next
															  /		  /
	tail		D									C
								- next ->




*/
