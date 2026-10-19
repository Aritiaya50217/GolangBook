package main

import "fmt"

// Node struct
type Node struct {
	nextNode *Node
	property rune
}

// create List method
func CreateLinkedList() *Node {
	headNode := &Node{nil, 'a'}
	currNode := headNode
	for i := 'b'; i <= 'z'; i++ {
		node := &Node{nil, i}
		currNode.nextNode = node
		currNode = node
	}
	return headNode
}

// Stringify  List method
func StringifyList(nodeList *Node) {
	currNode := nodeList
	for {
		fmt.Printf("%c", currNode.property)

		if currNode.nextNode != nil {
			currNode = currNode.nextNode
		} else {
			break
		}
	}
	fmt.Println("")
}

func ReverseLinkedList(nodeList *Node) *Node {
	currNode := nodeList
	var topNode *Node

	for {
		if currNode == nil {
			break
		}

		tempNode := currNode.nextNode
		currNode.nextNode = topNode
		topNode = currNode
		currNode = tempNode
	}

	return topNode
}

func main() {
	linkedList := CreateLinkedList()
	StringifyList(linkedList)
	StringifyList(ReverseLinkedList(linkedList))
}
