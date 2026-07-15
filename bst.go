package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func treeInsert(node *treeNode, val int) *treeNode {
	if node == nil {
		return &treeNode{
			val: val,
		}
	}
	if val < node.val {
		node.left = treeInsert(node.left, val)
	} else {
		node.right = treeInsert(node.right, val)
	}
	return node
}

func treeDelete(node *treeNode, val int) *treeNode {
	if node == nil {
		return node
	}
	if val < node.val {
		node.left = treeDelete(node.left, val)
		return node
	}
	if val > node.val {
		node.right = treeDelete(node.right, val)
		return node
	}
	if node.left == nil {
		return node.right
	}
	if node.right == nil {
		return node.left
	}
	successor := getSuccessor(node)
	node.val = successor.val
	node.right = treeDelete(node.right, successor.val)
	return node
}

func getSuccessor(node *treeNode) *treeNode {
	successor := node.right
	for successor != nil && successor.left != nil {
		successor = successor.left
	}
	return successor
}

func printTree(node *treeNode) {
	if node == nil {
		return
	}
	printTree(node.left)
	fmt.Printf("%d ", node.val)
	printTree(node.right)
}

func main() {
	var root *treeNode
	root = treeInsert(root, 10)
	root = treeInsert(root, 7)
	root = treeInsert(root, 12)
	root = treeInsert(root, 8)
	root = treeInsert(root, 5)
	root = treeInsert(root, 11)
	root = treeInsert(root, 15)

	printTree(root) // 5 7 8 10 11 12 15
	fmt.Println()

	root = treeDelete(root, 8)

	printTree(root) // 5 7 10 11 12 15
	fmt.Println()

	root = treeDelete(root, 15)

	printTree(root) // 5 7 10 11 12
	fmt.Println()
}
