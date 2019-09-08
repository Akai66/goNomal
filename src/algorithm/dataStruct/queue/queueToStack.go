package main

import (
	"algorithm/dataStruct/queue/qdata"
	"fmt"
)

func main() {
	ms := &myStack{new(qdata.MyQueue), new(qdata.MyQueue)}
	ms.Push("a")
	ms.Push("b")
	fmt.Printf("size:%d\n", ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	fmt.Printf("pop value:%v,size:%d\n", ms.Pop(), ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	ms.Push("c")
	ms.Push("d")
	fmt.Printf("size:%d\n", ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	fmt.Printf("pop value:%v,size:%d\n", ms.Pop(), ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	fmt.Printf("pop value:%v,size:%d\n", ms.Pop(), ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	ms.Push("e")
	fmt.Printf("size:%d\n", ms.Size())
	fmt.Printf("pop value:%v,size:%d\n", ms.Pop(), ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	fmt.Printf("isEmpty:%v\n", ms.Empty())
	fmt.Printf("pop value:%v,size:%d\n", ms.Pop(), ms.Size())
	fmt.Printf("Peek:%v\n", ms.Peek())
	fmt.Printf("isEmpty:%v\n", ms.Empty())
}

/*
用两个队列实现栈
- 用两个队列 queue 队列 与 help 队列
- Push 时直接 push 进 queue队列
- Pop 时先检查queue队列是否为空，若为空，说明所要实现的栈中不存在元素。若不为空，则将 queue队列中的元素装入 help中，queue中只剩最后一个，那么这便是应该Pop的元素，将该元素pop。pop之后，交换两队列，原help队列作现queue，原queue作现help即可。
*/

type myStack struct {
	queue *qdata.MyQueue
	helpq *qdata.MyQueue
}

func (ms *myStack) Push(v interface{}) {
	ms.queue.Push(v)
}

func (ms *myStack) Pop() interface{} {
	if ms.queue.Size() > 0 {
		for ms.queue.Size() > 1 {
			ms.helpq.Push(ms.queue.Pop())
		}
		value := ms.queue.Pop()
		ms.swap()
		return value
	}
	return nil
}

func (ms *myStack) Peek() interface{} {
	if ms.queue.Size() > 0 {
		for ms.queue.Size() > 1 {
			ms.helpq.Push(ms.queue.Pop())
		}
		value := ms.queue.Pop()
		ms.helpq.Push(value)
		ms.swap()
		return value
	}
	return nil
}

//交换
func (ms *myStack) swap() {
	ms.queue, ms.helpq = ms.helpq, ms.queue
}

func (ms *myStack) Size() int {
	return ms.queue.Size()
}

func (ms *myStack) Empty() bool {
	if ms.queue.Size() > 0 {
		return false
	}
	return true
}
