package datastruct

import "fmt"

type LkNode struct {
	Value interface{}
	Next *LkNode
}

func PrintLink(head *LkNode)  {
	for node := head;node != nil;node = node.Next {
		if node.Next != nil {
			fmt.Print(node.Value,"-->")
		}else {
			fmt.Println(node.Value)
		}
	}
}
