package main

import (
	"fmt"
	"sync"
)

// TreeNode class
type TreeNode struct {
	key       int
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// BinarySearchTree class
type BinarySearchTree struct {
	rootNode *TreeNode
	lock     sync.RWMutex
}

// InsertElement method
func (tree *BinarySearchTree) InsertElement(key, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	treeNode := TreeNode{
		key:       key,
		value:     value,
		leftNode:  nil,
		rightNode: nil,
	}
	if tree.rootNode == nil {
		tree.rootNode = &treeNode
	} else {
		insertTreeNode(tree.rootNode, &treeNode)
	}
}

// insertTreeNode
func insertTreeNode(rootNode, newTreeNode *TreeNode) {
	if newTreeNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newTreeNode
		} else {
			insertTreeNode(rootNode.leftNode, newTreeNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newTreeNode
		} else {
			insertTreeNode(rootNode.rightNode, newTreeNode)
		}
	}
}

// RemoveNode method
func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, key)
}

// removeNode method
func removeNode(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if key < treeNode.key {
		treeNode.leftNode = removeNode(treeNode.leftNode, key)
		return treeNode
	}

	if key > treeNode.key {
		treeNode.rightNode = removeNode(treeNode.rightNode, key)
		return treeNode
	}

	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}

	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}

	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}

	var leftmostrightNode *TreeNode
	leftmostrightNode = treeNode.rightNode
	for {
		// find smallest value on the right side
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.key, treeNode.value = leftmostrightNode.key, leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.key)
	return treeNode
}

// InOrderTraverse
func (tree *BinarySearchTree) InOrderTraverse(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

// inOrderTraverseTree method
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.rootNode, function)
}

// PreOrderTraverse method
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.rootNode, function)
}

// preOrderTraverseTree method
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// SearchNode method
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

// searchNode method
func searchNode(treeNode *TreeNode, key int) bool {
	if treeNode == nil {
		return false
	}
	if key < treeNode.key {
		return searchNode(treeNode.leftNode, key)
	}

	if key > treeNode.key {
		return searchNode(treeNode.rightNode, key)
	}
	return true
}

// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("********************************")
	stringify(tree.rootNode, 0)
	fmt.Println("********************************")
}

// stringify method
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "*** > "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.key)
		stringify(treeNode.rightNode, level)
	}
}

type TreeSet struct {
	bst *BinarySearchTree
}

func (treeset *TreeSet) InsertTreeNode(treeNodes ...TreeNode) {
	var treeNode TreeNode

	for _, treeNode = range treeNodes {
		treeset.bst.InsertElement(treeNode.key, treeNode.value)
	}
}

func (treeset *TreeSet) Delete(treeNodes ...TreeNode) {
	for _, treeNode := range treeNodes {
		treeset.bst.RemoveNode(treeNode.key)
	}
}

// Search method
func (treeset *TreeSet) Search(treeNodes ...TreeNode) bool {
	var treeNode TreeNode
	var exists bool
	for _, treeNode = range treeNodes {
		if exists = treeset.bst.SearchNode(treeNode.key); !exists {
			return false
		}
	}
	return true
}

func (treeset *TreeSet) String() {
	treeset.bst.String()
}

func main() {
	var treeset *TreeSet = &TreeSet{}
	treeset.bst = &BinarySearchTree{}
	var node1 TreeNode = TreeNode{8, 8, nil, nil}
	var node2 TreeNode = TreeNode{3, 3, nil, nil}
	var node3 TreeNode = TreeNode{10, 10, nil, nil}
	var node4 TreeNode = TreeNode{1, 1, nil, nil}
	var node5 TreeNode = TreeNode{6, 6, nil, nil}
	treeset.InsertTreeNode(node1, node2, node3, node4, node5)
	treeset.String()
}
