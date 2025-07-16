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
	head.data = "San Francisco"
	head.prev = nil
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", prev: head, next: nil}
	head.next = nodeOakland
	var nodeBerkeley *Node = &Node{data: "Berkeley", prev: nodeOakland, next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Fremont"
	tail.prev = nodeBerkeley
	tail.next = nil
	nodeBerkeley.next = tail
}

func removeNode(removePosition int) {
	p := head
	i := 0
	// Move the node to the previous node position the was deleted
	for {
		if p.next == nil || i >= removePosition-1 {
			break
		}
		p = p.next
		i++
	}
	// save the node you want to delete
	temp := p.next
	// Previous node next points to next of delete the node
	p.next = p.next.next
	p.next.prev = p
	// set the delete node next to null
	temp.next = nil
	// set the delete node prev to null
	temp.prev = nil
}

func output(node *Node) {
	p := node
	var end *Node = nil
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s-> ", p.data)
		end = p
		p = p.next
	}
	fmt.Printf("End\n")

	p = end
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s-> ", p.data)
		p = p.prev
	}
	fmt.Printf("Start\n\n")
}
func main() {
	initial()
	fmt.Printf("Delete a new node Berkeley at index = 2 :\n")
	removeNode(2)
	output(head)
}

/*	delete the index=2 node
		head					p											tail

					-->						-->					-->						-->
	San Francisco			Oakland					Berkeley				Fremont
					<--						<--					<--						<--

		head					p					temp					tail
												Node temp = p.next
					-->						-->					-->						-->
	San Francisco			Oakland					Berkeley				Fremont
					<--						<--					<--						<--

		head					p					temp					tail
						p.next = p.next.next
					-->											-->						-->
	San Francisco			Oakland					Berkeley				Fremont
					<--						<--					<--						<--

		head					p					temp					tail
												p.next.prev = p
					-->											-->						-->
	San Francisco			Oakland					Berkeley				Fremont
					<--						<--										<--



	After	temp.next	= nil and temp.prev = nil

		head					p					tail

					-->						-->
	San Francisco			Oakland					Fremont
					<--						<--

*/
