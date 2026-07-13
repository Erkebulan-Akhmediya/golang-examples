package main

import "fmt"

type node struct {
	val  int
	next *node
}

func preprend(head *node, val int) *node {
	n := node{
		val: val,
	}

	if head == nil {
		return &n
	}

	n.next = head
	return &n
}

func append(head *node, val int) {
	currentNode := head
	for currentNode != nil && currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = &node{val: val}
}

func printList(head *node) {
	currentNode := head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.val)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	var head *node
	head = preprend(head, 1)
	head = preprend(head, 1)
	head = preprend(head, 2)
	head = preprend(head, 3)

	printList(head) // [3 2 1 1]

	append(head, 6)
	append(head, 7)
	append(head, 8)

	printList(head) // [3 2 1 1 6 7 8]
}
