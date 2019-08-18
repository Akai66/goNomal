package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

//栈的特点：先进后出，后进先出
type Stack struct {
	head *Node
	size int
}

func main() {
	s := new(Stack)
	fmt.Println(s.empty())
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.pop()
	s.printStack()
	fmt.Println(s.empty(), s.getSize())
}

func (s *Stack) printStack() {
	head := s.head
	for head != nil {
		fmt.Print(head.value)
		head = head.next
	}
	fmt.Println()
}

func (s *Stack) pop() {
	if s.size > 0 {
		next := s.head.next
		s.head.next = nil
		s.head = next
		s.size -= 1
	}
}

func (s *Stack) push(v int) {
	//头插法
	newHead := new(Node)
	newHead.value = v
	if s.head == nil {
		s.head = newHead
	} else {
		newHead.next = s.head
		s.head = newHead
	}
	s.size += 1
}

func (s *Stack) empty() bool {
	if s.size >= 1 {
		return false
	} else {
		return true
	}
}

func (s *Stack) getSize() int {
	return s.size
}
