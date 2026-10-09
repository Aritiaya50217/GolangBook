package main

import (
	"fmt"
)

// Node class
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// IterateList method of LinkedList
func (linkedList *LinkedList) IterateList() {
	var node *Node
	properties := []int{}
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		properties = append(properties, node.property)
	}
	fmt.Println("properties : ", properties)
}

// NodeBetweenValue method of LinkedList
func (linkedList *LinkedList) NodeBetweenValue(firstProperty, secondProperty int) *Node {
	var nodeWith *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead method of LinkedList
func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

// NodeWithValue method of LinkedList
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var nodeWith *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter method of LinkedList
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

// LastNode method of LinkedList
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method of LinkedList
func (linkedList *LinkedList) AddToEnd(property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func main() {
	linkedList := LinkedList{}
	// headNode
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)

	// endNode
	linkedList.AddToEnd(5)
	// add after
	linkedList.AddAfter(1, 7)
	fmt.Println("headNode : ", linkedList.headNode.property)
	// all
	linkedList.IterateList()
	node := linkedList.NodeBetweenValue(1, 5)
	fmt.Println("NodeBetweenValue 1 and 5: ", node.property)

}
