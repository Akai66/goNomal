package main

import (
	"algorithm/dataStruct/queue/qdata"
	"algorithm/dataStruct/stack/data"
	"fmt"
)

//颠倒队列的前k个元素
//1.先将队列的前k个元素依次入栈
//2.将栈中的元素依次出栈，加入原队列
//3.最后将原队列的前size-k个元素出队列再入队列
func reverseKQueue(k int, q *qdata.MyQueue) error {
	if k <= 0 || k > q.Size() {
		return fmt.Errorf("参数错误")
	}
	qsize := q.Size()
	ms := new(data.MyStack)
	for i := 1; i <= k; i++ {
		ms.Push(q.Pop())
	}
	for !ms.Empty() {
		q.Push(ms.Pop())
	}
	for j := 1; j <= qsize-k; j++ {
		q.Push(q.Pop())
	}
	return nil
}

func main() {
	q := new(qdata.MyQueue)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.Push(6)
	q.Print()
	err := reverseKQueue(3, q)
	if err != nil {
		fmt.Println(err)
	}
	q.Print()
}
