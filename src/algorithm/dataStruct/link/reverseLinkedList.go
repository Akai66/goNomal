package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func main() {
	head := new(Node)
	//new(Node)返回的是Node结构体的指针，指针访问结构体成员字段的方法是(*head).value=1，但是go语言底层做了优化，会自动把head.value转换成(*head).value
	head.value = 1
	node1 := new(Node)
	node1.value = 2
	head.next = node1
	node2 := new(Node)
	node2.value = 3
	node1.next = node2
	printLinkedList(head)
	head = reverseLinkedList(head)
	printLinkedList(head)
}

func printLinkedList(head *Node) {
	for ; head != nil; head = head.next {
		fmt.Print(head.value)
	}
	fmt.Println()
}

//头插法
func reverseLinkedList(head *Node) *Node {
	newHead := new(Node)
	for head != nil {
		next := head.next
		head.next = newHead.next
		newHead.next = head
		head = next
	}
	return newHead.next
}
