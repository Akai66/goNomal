package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	node1 := new(Node)
	node1.value = 9
	node2 := new(Node)
	node2.value = 9
	node1.next = node2
	node3 := new(Node)
	node3.value = 2
	node2.next = node3
	node4 := new(Node)
	node4.value = 9
	node3.next = node4

	onode1 := new(Node)
	onode1.value = 8
	onode2 := new(Node)
	onode2.value = 1
	onode1.next = onode2

	head := addTowLink(reverseLink(node1), reverseLink(onode1))
	printLink(head)
}

func printLink(head *Node) {
	for ; head != nil; head = head.next {
		fmt.Print(head.value)
	}
	fmt.Println()
}

//反转链表
func reverseLink(head *Node) *Node {
	newHead := new(Node)
	for head != nil {
		next := head.next
		head.next = newHead.next
		newHead.next = head
		head = next
	}
	return newHead.next
}

//相加后采用头插法
func addTowLink(head1, head2 *Node) *Node {
	newHead := new(Node)
	n1, n2 := head1, head2
	var carry int
	for n1 != nil || n2 != nil || carry > 0 {
		var sum, v int
		if n1 != nil {
			sum += n1.value
			n1 = n1.next
		}
		if n2 != nil {
			sum += n2.value
			n2 = n2.next
		}
		sum += carry
		v = sum % 10
		carry = sum / 10
		node := new(Node)
		node.value = v
		node.next = newHead.next
		newHead.next = node
	}
	return newHead.next
}
