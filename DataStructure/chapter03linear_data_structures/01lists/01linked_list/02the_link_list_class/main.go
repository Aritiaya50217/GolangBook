package main

import "fmt"

// node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// AddToHead method of LinkedList class
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = node
}

/*
		IterateList method iterates over LinkedList
	 	IterateList คือ วิธีการวนลูปใน LinkedList ใน Singly LinkedList  สามารถทำได้โดยการใช้ลูป (for)
		เพื่อวนไปที่แต่ละ node ใน LinkedList ตั้งแต่ node แรกจนถึง node สุดท้าย
*/
func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println("IterateList : ", node.property)
	}
}

// LastNode method returns the last Node
func (linkedList *LinkedList) LastNode() *Node {
	var lastNode *Node
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd method adds the node with property to the end
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{
		property: property,
		nextNode: nil,
	}
	lastNode := linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// NodeWithValue method returns Node given parameter property
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

// AddAfter method adds a node with nodeProperty after node with property
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}
	nodeWith := linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		fmt.Println("node.nextNode.property : ", node.nextNode.property)
		fmt.Println("nodeWith.nextNode.property : ", nodeWith.nextNode.property)
		nodeWith.nextNode = node // เอา node ใหม่มาแทรก
	}

}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Println("AddToHead : ", linkedList.headNode)
	linkedList.IterateList()
	linkedList.LastNode()
	fmt.Println("LastNode : ", linkedList.headNode)
	linkedList.AddToEnd(4)
	linkedList.IterateList()
	fmt.Println("After AddToEnd : ", linkedList)
	linkedList.NodeWithValue(4)
	fmt.Println("After NodeWithValue : ", linkedList.headNode.nextNode.property)
	linkedList.AddAfter(1, 7)
	linkedList.IterateList()
}
