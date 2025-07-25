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

func output() {
	var p = head
	for {
		fmt.Printf("%s-> ", p.data)
		p = p.next
		// วนครบ loop
		if p == head {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("End\n")

	p = tail
	for {
		fmt.Printf("%s-> ", p.data)
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
	output()
}

/*
				p
								p = p.next
	  		   head

								- next ->
				A	 								B
			|		|			- prev ->		|		|
p = p.next		next	prev						prev	next 		p = p.next
			|		|							|		|
								- prev ->
	tail		D									C
								- next ->

								p = p.next





									p = p.prev
				head
									- next ->
					A	 								B
				|		|			- prev ->		|		|
p = p.prev		next	prev						prev	next 		p = p.prev
				|		|							|		|
									- prev ->
p		tail		D									C
									- next ->

			p = p.prev

*/
