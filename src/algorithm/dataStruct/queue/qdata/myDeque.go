package qdata

import "fmt"

type qNode struct {
	pre   *qNode
	value interface{}
	next  *qNode
}

type MyDeque struct {
	head *qNode
	tail *qNode
	size int
}

func (mdq *MyDeque) Push(v interface{}) {
	node := new(qNode)
	node.value = v
	if mdq.size > 0 {
		mdq.tail.next = node
		node.pre = mdq.tail
		mdq.tail = node
	} else {
		mdq.head = node
		mdq.tail = node
	}
	mdq.size += 1
}

func (mdq *MyDeque) Pop() interface{} {
	if mdq.size > 0 {
		next := mdq.head.next
		value := mdq.head.value
		mdq.head.next = nil
		mdq.head = next
		if next != nil {
			next.pre = nil
		} else {
			mdq.tail = nil
		}
		mdq.size -= 1
		return value
	}
	return nil
}

func (mdq *MyDeque) Peek() interface{} {
	if mdq.size > 0 {
		return mdq.head.value
	}
	return nil
}

func (mdq *MyDeque) GetLast() interface{} {
	if mdq.size > 0 {
		return mdq.tail.value
	}
	return nil
}

func (mdq *MyDeque) RemoveLast() interface{} {
	if mdq.size > 0 {
		pre := mdq.tail.pre
		value := mdq.tail.value
		mdq.tail.pre = nil
		mdq.tail = pre
		if pre != nil {
			pre.next = nil
		} else {
			mdq.head = nil
		}
		mdq.size -= 1
		return value
	}
	return nil
}

func (mdq *MyDeque) Print() {
	head := mdq.head
	for head != nil {
		if head.next != nil {
			fmt.Printf("%v->", head.value)
		} else {
			fmt.Println(head.value)
		}
		head = head.next
	}
}

func (mdq *MyDeque) PrintReverse() {
	tail := mdq.tail
	for tail != nil {
		if tail.pre != nil {
			fmt.Printf("%v->", tail.value)
		} else {
			fmt.Println(tail.value)
		}
		tail = tail.pre
	}
}

func (mdq *MyDeque) Size() int {
	return mdq.size
}

func (mdq *MyDeque) Empty() bool {
	if mdq.size > 0 {
		return false
	}
	return true
}
