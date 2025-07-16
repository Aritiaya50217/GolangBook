package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.next = nil

	var nodeB *Node = &Node{data: "B", next: nil}
	head.next = nodeB

	var nodeC *Node = &Node{data: "C", next: nil}
	nodeB.next = nodeC

	tail.data = "D"
	// วนกลับไปที่ head เพราะเป็น circle
	tail.next = head
	nodeC.next = tail

}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0
	// move the node to the insertion postion
	for {
		if p.next == nil || i >= insertPosition {
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
}

func output(node *Node) {
	var p = node
	for {
		fmt.Printf("%s-> ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s\n\n", p.data)
}

func main() {
	initial()
	fmt.Printf("Insert a new node E at index = 2 : \n")
	insert(2, "E")
	output(head)
}

/*	Insert a node E in position 2

				p

	  		   head						 newNode
				A	- next -> 	B			E
				|				|
 		      next			   next
	   			|				|
	tail   		D	<- next -	C


					 p=p.next
								p
	  		   head
				A	- next -> 	B
				|				|
 		      next			   next
	   			|				|
	tail   		D	<- next -	C


					   newNode.next = p.next
	  		   head								p
				A			- next -> 			B
				|								|
 		      next			  			 	   next		 E
			  											/
	   			|								|	next
													/
	tail   		D			<- next -			C


					   	p.next = newNode
	  		   head								p
				A			- next -> 			B
													\
													next
														\
				|								|
 		      next			  			 	   next		 E
			  											/
	   			|								|	next
													/
	tail   		D			<- next -			C


*/
