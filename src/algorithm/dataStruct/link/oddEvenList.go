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
	head := oddEvenList(node1)
	printLink(head)
}

func printLink(head *Node) {
	for ; head != nil; head = head.next {
		fmt.Print(head.value)
	}
	fmt.Println()
}

//Example:
//Given 1->2->3->4->5->NULL,
//return 1->3->5->2->4->NULL.
func oddEvenList(head *Node) *Node {
	if head == nil {
		return head
	}
	odd, even, evenHead := head, head.next, head.next
	for even != nil && even.next != nil {
		odd.next = odd.next.next
		odd = odd.next
		even.next = even.next.next
		even = even.next
	}
	odd.next = evenHead
	return head
}
