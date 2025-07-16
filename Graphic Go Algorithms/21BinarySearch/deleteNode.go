package main

import (
	"fmt"
)

/*	binary search tree Delete Node
	- If there are two child nodes , replace the current node with the smallest
	node from the right subtree , Delete node 40

		    				  root
						   	   60
							/  	   \2
		delete  40		  40		80
			 			/	\  	   /  \
				 	  30     50	 70	   90


					  	      root
						   	   60
							/  	   \2
						  50		80
			 			/	  	   /  \
				 	  30     	 70	   90

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

func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf("%d ", root.data)
	inOrder(root.right)
}

func removeNode(node *Node, newData int) *Node {
	if node == nil {
		return node
	}
	var compareValue = newData - node.data
	if compareValue > 0 {
		node.right = removeNode(node.right, newData)
	} else if compareValue < 0 {
		node.left = removeNode(node.left, newData)
	} else if node.left != nil && node.right != nil {
		// find the minimum node of the right subtree to replace the current node
		node.data = searchMinValue(node.right).data
		node.right = removeNode(node.right, node.data)
	} else {
		if node.left != nil {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
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
	//Constructing a binary search tree
	insert(root, 60)
	insert(root, 40)
	insert(root, 20)
	insert(root, 10)
	insert(root, 30)
	insert(root, 50)
	insert(root, 80)
	insert(root, 70)
	insert(root, 90)
	fmt.Printf("\ndelete node is: 10 \n")
	removeNode(root, 10)
	fmt.Printf("\nIn-order traversal binary tree \n")
	inOrder(root)
	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("\ndelete node is: 20 \n")
	removeNode(root, 20)
	fmt.Printf("\nIn-order traversal binary tree \n")
	inOrder(root)
	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("\ndelete node is: 40 \n")
	removeNode(root, 40)
	fmt.Printf("\nIn-order traversal binary tree \n")
	inOrder(root)

}
