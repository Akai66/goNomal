package main

import (
	"algorithm/swordOffer/datastruct"
	"fmt"
)

func main()  {
	node1 := new(datastruct.LkNode)
	node1.Value = 1
	node2 := new(datastruct.LkNode)
	node2.Value = 2
	node1.Next = node2
	node3 := new(datastruct.LkNode)
	node3.Value = 3
	node2.Next = node3
	node4 := new(datastruct.LkNode)
	node4.Value = 3
	node3.Next = node4
	node5 := new(datastruct.LkNode)
	node5.Value = 4
	node4.Next = node5
	node6 := new(datastruct.LkNode)
	node6.Value = 5
	node5.Next = node6
	node7 := new(datastruct.LkNode)
	node7.Value = 5
	node6.Next = node7
	linkDelDuplication(&node1)
	datastruct.PrintLink(node1)
}

func linkDelDuplication(head **datastruct.LkNode) error{
	if head == nil || *head == nil {
		fmt.Errorf("参数错误:头结点为nil")
	}
	var preNode *datastruct.LkNode = nil
	curNode := *head

	for curNode != nil {
		isdel := false
		next := curNode.Next
		if next != nil && next.Value == curNode.Value {
			isdel = true
		}
		if !isdel {
			preNode = curNode
			curNode = next
		}else {
			value := curNode.Value
			tobeDelNode := curNode
			for tobeDelNode != nil && tobeDelNode.Value == value {
				next = tobeDelNode.Next
				tobeDelNode = next
			}
			if preNode == nil {
				*head = next
			}else{
				preNode.Next = next
			}
			curNode = next
		}
	}
	return nil
}