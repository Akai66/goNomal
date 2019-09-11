package qdata

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

//链表实现队列
type MyQueue struct {
	head *Node
	size int
}

//入队列
func (mq *MyQueue) Push(v interface{}) {
	node := new(Node)
	node.value = v
	tail := mq.head
	if tail == nil {
		mq.head = node
	} else {
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = node
	}
	mq.size += 1
}

//出队列
func (mq *MyQueue) Pop() interface{} {
	if mq.size > 0 {
		next := mq.head.next
		value := mq.head.value
		mq.head.next = nil
		mq.head = next
		mq.size -= 1
		return value
	}
	return nil
}

func (mq *MyQueue) Top() interface{} {
	if mq.size > 0 {
		return mq.head.value
	}
	return nil
}

func (mq *MyQueue) Empty() bool {
	if mq.size > 0 {
		return false
	}
	return true
}

func (mq *MyQueue) Size() int {
	return mq.size
}

func (mq *MyQueue) Print() {
	head := mq.head
	for head != nil {
		if head.next != nil {
			fmt.Printf("%v->", head.value)
		} else {
			fmt.Println(head.value)
		}
		head = head.next
	}
}
