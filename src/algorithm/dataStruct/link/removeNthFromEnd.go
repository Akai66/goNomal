package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	node1 := new(Node)
	node1.value = 1
	node2 := new(Node)
	node2.value = 2
	node1.next = node2
	node3 := new(Node)
	node3.value = 3
	node2.next = node3
	node4 := new(Node)
	node4.value = 4
	node3.next = node4
	node5 := new(Node)
	node5.value = 5
	node4.next = node5
	node6 := new(Node)
	node6.value = 6
	node5.next = node6
	printLink(node1)
	head := removeNthFromEnd(node1, 6)
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

func removeNthFromEnd(head *Node, n int) *Node {
	fast := head
	for ; n > 0; n-- {
		fast = fast.next
	}
	if fast == nil {
		return head.next
	}
	low := head
	for fast.next != nil {
		fast = fast.next
		low = low.next
	}
	low.next = low.next.next
	return head
}
