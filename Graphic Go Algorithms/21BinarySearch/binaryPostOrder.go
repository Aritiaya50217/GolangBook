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

/*	binary search tree Post-order traversal
	Post-order traversal : left subtree -> right subtree -> root node

	Result :  [	10 30 20 50 40 80 70 90 60]
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

// Post-Order
func postOrder(root *Node) {
	if root == nil {
		return
	}
	// Recursive Traversing the left subtree
	postOrder(root.left)
	// Recursive Traversing the right subtree
	postOrder(root.right)
	fmt.Printf("%d ", root.data)

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
	fmt.Printf("Post-order traversal binary search tree \n")
	postOrder(root)
}
