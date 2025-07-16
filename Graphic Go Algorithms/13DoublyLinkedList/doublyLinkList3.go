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
	newNode.next.prev = newNode
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
	fmt.Printf("Insert a new node Walnut at index 2 : \n")
	insert(2, "Walnut")
	output(head)
}

/*	Insert a node Walnut in position 2
		head					p					newNode						tail
													Walnut
					-->						-->					-->
	San Francisco			Oakland					Berkeley					Fremont
					<--						<--					<--

		head					p				newNode.next = p.next						tail
													Walnut
					-->						-->					-->				-->
	San Francisco			Oakland								Berkeley					Fremont
					<--						<--					<--				<--

		head					p				p.next = newNode 							tail
													Walnut
					-->						-->					-->					-->
	San Francisco			Oakland								Berkeley					Fremont
					<--						<--					<--					<--

		head					p				newNode.prev = p							tail
													Walnut
					-->						-->					-->					-->
	San Francisco			Oakland								Berkeley					Fremont
					<--						<--					<--					<--

		head					p																	tail
											newNode.next.prev = newNodd
					-->						-->					-->						-->
	San Francisco			Oakland					Walnut					Berkeley				Fremont
					<--						<--					<--						<--

*/
