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
	tail.next = head
	nodeC.next = tail
}

func removeNode(removePosition int) {
	var p = head
	var i = 0
	// move the node to the previous node postion that was deleted
	for {
		//  removePosition-1 => ลำดับถัดไป
		if p.next == nil || i >= removePosition-1 {
			break
		}
		p = p.next
		i++
	}
	// save the node you want to delete
	var temp = p.next
	// previous node next points to next of delete the node
	p.next = p.next.next
	temp.next = nil
}

func output(node *Node) {
	p := node
	for {
		fmt.Printf("%s-> ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s \n\n", p.data)

}

func main() {
	initial()
	fmt.Printf("Delete a new node E at index = 2 :\n")
	// ตำแหน่งที่ถูกลบ
	removeNode(2)
	output(head)
}

/*
				p

	  		   head
				A	- next -> 	B
				|				|
 		      next			   next
	   			|				|
	tail   		D	<- next -	C

					p = p.next
	  		   head				p
				A	- next -> 	B
				|				|
 		      next			   next
	   			|				|
	tail   		D	<- next -	C


					Node temp = p.next
	  		   head						p
				A		- next -> 		B
				|						|
 		      next			   		  next
	   			|						|
	tail   		D		<- next -		C	--  temp

				p.next=p.next.next
	  		   head		  p
				A- next ->B
				|	     /
 		      next	 next
	   			|	/
	tail   		D<- next- C	--  temp

				temp.next = nil
	  		   head		  p
				A- next ->B
				|	     /
 		      next	 next
	   			|	/
	tail   		D		  C	--  temp (โดนลบ)


*/
