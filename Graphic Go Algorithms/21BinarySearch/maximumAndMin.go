package main

import (
	"fmt"
)

/*	Construct a binary search tree , insert node

	Insert 60
					root
					 60

	Insert 40
					40 < 60
					  root
					   60
					/
				40

	Insert 20
					  		root
			20 < 60		   	 60
							/
		20 < 40			 40
			 			/
		 			20

	Insert 10
							root
			10 < 60		   	 60
							/
		10 < 40			 40
			 			/
	10 < 20 	 	  20
				    /
			     10


	Insert 30
							root
			30 < 60		   	 60
							/
		30 < 40			 40
			 			/
				 	  20     30 > 20
				    /	\
			     10 	  30

	.
	.
	.

	Finished
							  root
						   	   60
							/  	   \2
			  			  40		80
			 			/	\  	   /  \
				 	  20     50	 70	   90
				    /	\
			     10 	  30

*/

/*	binary search tree Maximum and minimum
	Minimum value : ค่าที่น้อยที่สุด ดูจากฝั่งซ้ายของ root node
	Maximum value : ค่าที่มากที่สุด ดูจากฝั่งขวาของ root node
*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

var root *Node = nil

func createNewNode(newData int) *Node {
	var newNode *Node = new(Node)
	newNode.data = newData
	newNode.left = nil
	newNode.right = nil
	return newNode
}

func searchMinValue(node *Node) *Node {
	if node == nil || node.data == 0 {
		return nil
	}
	if node.left == nil {
		return node
	}
	// recursively find the minimum from the left subtree
	return searchMinValue(node.left)
}

func searchMaxValue(node *Node) *Node {
	if node == nil || node.data == 0 {
		return nil
	}
	if node.right == nil {
		return node
	}
	return searchMaxValue(node.right)
}

func insert(node *Node, newData int) {
	if root == nil {
		root = &Node{data: newData, left: nil, right: nil}
		return
	}
	var compareValue = newData - node.data
	// recursive left subtree , continue to find the insertion position
	if compareValue < 0 {
		if node.left == nil {
			node.left = createNewNode(newData)
		} else {
			insert(node.left, newData)
		}
	} else if compareValue > 0 {
		if node.right == nil {
			node.right = createNewNode(newData)
		} else {
			insert(node.right, newData)
		}
	}
}

func main() {
	// constructing a binary search tree
	insert(root, 60)
	insert(root, 40)
	insert(root, 90)
	insert(root, 20)
	fmt.Printf("\nMinimum Value \n")
	var minNode = searchMinValue(root)
	fmt.Printf("%d ", minNode.data)

	fmt.Printf("\nMaximum Value \n")
	var maxNode = searchMaxValue(root)
	fmt.Printf("%d ", maxNode.data)

}
