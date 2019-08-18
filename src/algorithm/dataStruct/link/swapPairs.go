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
	head := swapPairs(node1)
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

//Given 1->2->3->4, you should return the list as 2->1->4->3.
func swapPairs(head *Node) *Node {
	newHead := new(Node)
	newHead.next = head
	pre := newHead
	for pre.next != nil && pre.next.next != nil {
		n1 := pre.next
		n2 := pre.next.next
		next := n2.next
		n2.next = n1
		n1.next = next
		pre.next = n2
		pre = n1
	}
	return newHead.next
}
