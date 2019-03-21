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
	node2.value = 1
	node1.next = node2
	node3 := new(Node)
	node3.value = 2
	node2.next = node3
	node4 := new(Node)
	node4.value = 3
	node3.next = node4
	node5 := new(Node)
	node5.value = 3
	node4.next = node5
	node6 := new(Node)
	node6.value = 4
	node5.next = node6

	head := deleteDuplicates(node1)
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

//删除重复节点，前提是有序
//1-->1-->2-->3-->3-->4  =>  1-->2-->3-->4
func deleteDuplicates(head *Node) *Node {
	if head.next == nil || head == nil {
		return head
	}
	head.next = deleteDuplicates(head.next)
	if head.value == head.next.value {
		return head.next
	} else {
		return head
	}
}
