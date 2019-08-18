package main

import "fmt"

type Node struct {
	pre   *Node
	value int
	next  *Node
}

//队列的特点:先进先出
type Queue struct {
	head *Node
	tail *Node
	size int
}

func main() {
	q := new(Queue)
	q.add(1)
	q.add(2)
	q.add(3)
	q.add(4)
	q.remove()
	q.printQueue()
	q.remove()
	q.printQueue()
}

func (q *Queue) printQueue() {
	tail := q.tail
	for tail != nil {
		fmt.Print(tail.value)
		tail = tail.pre
	}
	fmt.Println()
}

func (q *Queue) add(v int) {
	newHead := new(Node)
	newHead.value = v
	if q.head == nil {
		q.head = newHead
		q.tail = newHead
	} else {
		newHead.next = q.head
		q.head.pre = newHead
		q.head = newHead
	}
	q.size++
}

func (q *Queue) remove() {
	if q.size > 0 {
		newTail := q.tail.pre
		if newTail == nil {
			q.head = nil
			q.tail = nil
		} else {
			newTail.next = nil
			q.tail = newTail
		}
		q.size--
	}
}
