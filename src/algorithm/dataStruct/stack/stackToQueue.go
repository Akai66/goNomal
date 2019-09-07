package main

import (
	"algorithm/dataStruct/stack/data"
	"fmt"
)

func main() {
	mq := &myQueue{new(data.MyStack), new(data.MyStack), 0}
	mq.Push("a")
	mq.Push("b")
	fmt.Printf("size:%d\n", mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	fmt.Printf("pop value:%v,size:%d\n", mq.Pop(), mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	mq.Push("c")
	mq.Push("d")
	fmt.Printf("size:%d\n", mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	fmt.Printf("pop value:%v,size:%d\n", mq.Pop(), mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	fmt.Printf("pop value:%v,size:%d\n", mq.Pop(), mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	mq.Push("e")
	fmt.Printf("size:%d\n", mq.Size())
	fmt.Printf("pop value:%v,size:%d\n", mq.Pop(), mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	fmt.Printf("isEmpty:%v\n", mq.Empty())
	fmt.Printf("pop value:%v,size:%d\n", mq.Pop(), mq.Size())
	fmt.Printf("top:%v\n", mq.Top())
	fmt.Printf("isEmpty:%v\n", mq.Empty())
}

//利用双栈实现队列，先进先出
type myQueue struct {
	instack  *data.MyStack
	outstack *data.MyStack
	size     int
}

//入队列
func (mq *myQueue) Push(v interface{}) {
	mq.instack.Push(v)
	mq.size += 1
}

//出队列
func (mq *myQueue) Pop() interface{} {
	if mq.outstack.Empty() {
		for !mq.instack.Empty() {
			mq.outstack.Push(mq.instack.Pop())
		}
	}
	if mq.outstack.Empty() {
		return nil
	}
	v := mq.outstack.Pop()
	mq.size -= 1
	return v
}

//获取队首元素
func (mq *myQueue) Top() interface{} {
	if mq.outstack.Empty() {
		for !mq.instack.Empty() {
			mq.outstack.Push(mq.instack.Pop())
		}
	}

	if mq.outstack.Empty() {
		return nil
	}
	return mq.outstack.Top()
}

//判断队列是否为空
func (mq *myQueue) Empty() bool {
	if mq.size > 0 {
		return false
	}
	return true
}

//获取队列元素个数
func (mq *myQueue) Size() int {
	return mq.size
}
