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
	headArr := splitListToParts(node1, 4)
	for _, head := range headArr {
		printLink(head)
	}
}

func printLink(head *Node) {
	for ; head != nil; head = head.next {
		fmt.Print(head.value)
	}
	fmt.Println()
}

//Input:
//root = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], k = 3
//Output: [[1, 2, 3, 4], [5, 6, 7], [8, 9, 10]]
//把链表分隔成 k 部分，每部分的长度都应该尽可能相同，排在前面的长度应该大于等于后面的
func splitListToParts(head *Node, k int) []*Node {
	var headArr []*Node
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.next
	}
	mod := length % k
	size := length / k
	cur = head
	for i := 0; cur != nil && i < k; i++ {
		headArr = append(headArr, cur)
		newSize := size
		if mod > 0 {
			newSize++
			mod--
		}
		for j := 0; j < newSize-1; j++ {
			cur = cur.next
		}
		next := cur.next
		cur.next = nil
		cur = next
	}
	return headArr
}
