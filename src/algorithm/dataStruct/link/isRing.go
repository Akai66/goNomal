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

	onode1 := new(Node)
	onode1.value = 8
	onode2 := new(Node)
	onode2.value = 6
	onode1.next = onode2
	onode3 := node3
	onode2.next = onode3
	node4.next = onode2

	fmt.Println(isRing(onode1))
}

//判断单链表是否有环
func isRing(head *Node) bool {
	if head == nil {
		return false
	}
	low := head.next
	if low == nil {
		return false
	}
	fast := low.next
	for low != nil && fast != nil {
		if low == fast {
			return true
		}
		low = low.next
		fast = fast.next
		if fast == nil {
			return false
		}
		fast = fast.next
	}
	return false
}
