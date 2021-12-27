package main

import "fmt"

type Node struct {
	data string
	next *Node
}

func PrintList(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	fmt.Println(node.next)
	PrintList(node.next)
}

//TODO add implementation for deleting node

func main() {
	node3 := &Node{data: "three"}
	node2 := &Node{data: "two", next: node3}
	node1 := &Node{data: "one", next: node2}
	PrintList(node1)
}
