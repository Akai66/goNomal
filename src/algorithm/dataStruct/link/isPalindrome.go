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
	node4.value = 2
	node3.next = node4
	node5 := new(Node)
	node5.value = 1
	node4.next = node5
	fmt.Println(isPalindrome(node1))
}

//判断是否为回文链表
func isPalindrome(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}
	low, fast := head, head.next
	for fast != nil && fast.next != nil {
		low = low.next
		fast = fast.next.next
	}
	if fast != nil {
		low = low.next //偶数节点，需要将low再往后走一步
	}
	//拆分链表
	cut(head, low)
	//反转后面的链表
	head2 := reverse(low)
	//比较两个链表是否相等
	return isEqual(head, head2)
}

func cut(head, low *Node) {
	for head.next != low {
		head = head.next
	}
	head.next = nil
}

func reverse(head *Node) *Node {
	newHead := new(Node)
	for head != nil {
		next := head.next
		head.next = newHead.next
		newHead.next = head
		head = next
	}
	return newHead.next
}

func isEqual(head1, head2 *Node) bool {
	for head1 != nil && head2 != nil {
		if head1.value != head2.value {
			return false
		}
		head1 = head1.next
		head2 = head2.next
	}
	return true
}
