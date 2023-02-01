package main

import "fmt"

type linkedList struct {
	next *node
}
type node struct {
	next  *node
	value string
}

func newLinkedList() *linkedList {
	return &linkedList{next: nil}
}
func (l *linkedList) insert(v string) {
	nodeToBeInserted := &node{value: v, next: l.next}
	l.next = nodeToBeInserted
}

func (l *linkedList) invert() {
	l.next = invertLinkedListRecursively(nil, l.next)
}

func invertLinkedListRecursively(currentNode, nextNode *node) *node {
	x := &node{}
	if nextNode.next != nil {
		x = invertLinkedListRecursively(nextNode, nextNode.next)
	} else {
		nextNode.next = currentNode
		return nextNode
	}
	nextNode.next = currentNode
	return x
}

func (l *linkedList) printList() {
	node := l.next
	for {
		if node == nil {
			break
		}
		fmt.Println(node)
		node = node.next
	}
}

func main() {
	names := newLinkedList()
	for i := 0; i < 100; i++ {
		names.insert("A")
	}
	for i := 0; i < 100000; i++ {
		names.invert()
	}
	fmt.Println("DONE!")

}
