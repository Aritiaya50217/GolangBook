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

func output(node *Node) {
	var p = node
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
	output(head)
}

/*	initialization

		head																tail
					-->						-->					-->					
	San Francisco			Oaklanda				Berkeley				Fremont
					<--						<--					<--				

*/
