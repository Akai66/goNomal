package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	node1 := new(Node)
	node1.value = 3
	node2 := new(Node)
	node2.value = 4
	node1.next = node2
	node3 := new(Node)
	node3.value = 7
	node2.next = node3
	node4 := new(Node)
	node4.value = 9
	node3.next = node4

	onode1 := new(Node)
	onode1.value = 5
	onode2 := new(Node)
	onode2.value = 12
	onode1.next = onode2

	head := mergeSortedLink(node1, onode1)
	printLink(head)
}

func printLink(head *Node) {
	for ; head != nil; head = head.next {
		if head.next != nil {
			fmt.Print(head.value, "-->")
		} else {
			fmt.Print(head.value)
		}
	}
	fmt.Println()
}

func mergeSortedLink(n1, n2 *Node) *Node {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.value < n2.value {
		n1.next = mergeSortedLink(n1.next, n2)
		return n1
	} else {
		n2.next = mergeSortedLink(n1, n2.next)
		return n2
	}
}
