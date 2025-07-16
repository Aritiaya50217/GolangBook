package main

import "fmt"

type Node struct {
	data string
	next *Node
}

// the first node called head node
var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "San Francisco"
	head.next = nil
	var nodeOakland *Node = &Node{data: "Oakland", next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", next: nil}
	nodeOakland.next = nodeBerkeley

	// var tail *Node = &Node{data: "Fremont", next: nil}
	// nodeBerkeley.next = tail

	tail.data = "Fremont"
	tail.next = nil
	nodeBerkeley.next = tail
}

func add(data string) {
	var newNode *Node = &Node{data: data, next: nil}
	tail.next = newNode
	tail = newNode
}

func output(node *Node) {
	p := node
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s->", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")

}
func main() {
	initial()
	add("Walnut")
	output(head)
}

/*
	head										tail
	San Francisco--> Oakland --> Berkeley --> Fremont

	Traversal output

		p
		head									tail
	San Francisco--> Oakland --> Berkeley --> Fremont

		head						p												tail
						p=head!=nil ,p=p.next=nodeOakland
	San Francisco--> 			Oakland --> 				Berkeley --> 			Fremont
	head														p					tail
												p=nodeOakland!=nil ,p=p.next=nodeBerkeley
	San Francisco--> 			Oakland --> 				Berkeley --> 			Fremont

	head																				p
																						tail
																			p=nodeBerkeley!=nil ,p=p.next=tail
	San Francisco--> 			Oakland --> 				Berkeley --> 			Fremont
																			p = tail != nil , p=p.next=nil , end

	Append a new node
	head
																					tail			create a node
	San Francisco--> 			Oakland --> 				Berkeley --> 			Fremont 			Walnut

	head
																					tail			tail.next = node.Walnut
	San Francisco--> 			Oakland --> 				Berkeley --> 			Fremont -->			Walnut
	head
																										tail
	San Francisco--> 			Oakland --> 				Berkeley --> 			Fremont -->			Walnut

*/
