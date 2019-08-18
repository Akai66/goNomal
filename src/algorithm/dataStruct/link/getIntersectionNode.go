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

	fmt.Println(getIntersectionNode(node1, onode1))
	fmt.Println(isIntersection1(node1, onode1))
	fmt.Println(isIntersection2(node1, onode1))

}

//获取两个链表的交点,前提是两个链表一定相交
func getIntersectionNode(head1, head2 *Node) int {
	n1, n2 := head1, head2
	for n1 != n2 {
		if n1.next != nil {
			n1 = n1.next
		} else {
			n1 = head2
		}
		if n2.next != nil {
			n2 = n2.next
		} else {
			n2 = head1
		}
	}
	return n1.value
}

//判断两个链表是否相交
//方法一：将链表1的尾节点连接到链表2的头节点，然后判断链表2是否有环
func isIntersection1(head1, head2 *Node) bool {
	n := head1
	for n.next != nil {
		n = n.next
	}
	//将链表1的尾节点连接到链表2的头节点
	n.next = head2
	//判断链表2是否有环
	if head2.next == nil {
		return false
	}
	low := head2.next
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

//方法二：判断两个链表最后一个节点是否相同
func isIntersection2(head1, head2 *Node) bool {
	//此处并不会改变原始参数
	for head1 != nil && head2 != nil {
		if head1 == head2 {
			return true
		} else {
			head1 = head1.next
			head2 = head2.next
		}
	}
	return false
}
