package main

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

func (l *LinkedList[T]) add(value T) {
	newNode := &Node[T]{value, nil}
	if l.head == nil {
		l.head = newNode
		l.size = 1
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	l.size++
}

func (l *LinkedList[T]) remove(value T) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.size--
		return
	}
	current := l.head
	for current.next != nil && current.next.data == value {
		current.next = current.next.next
		l.size--
		return
	}
}

func (l *LinkedList[T]) printList() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func main() {
	ll := LinkedList[int]{}
	ll.add(1)
	ll.add(2)
	ll.printList()
	fmt.Println("size:", ll.size)

	listString := LinkedList[string]{}
	listString.add("Hemnath")
	listString.add("Samantha")
	listString.add("Nayanthara")
	listString.printList()
	fmt.Println("size:", listString.size)
}
