package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func printList(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	printList(node.next)
}

func dummyList() {
	node3 := &node{data: 3}
	node2 := &node{data: 2, next: node3}
	node1 := &node{data: 1, next: node2}
	printList(node1)
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main() {
	// dummyList()
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 28}
	node3 := &node{data: 68}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	fmt.Println(myList)
	printList(myList.head)

}
