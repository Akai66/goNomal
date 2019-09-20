package main

import (
	"algorithm/dataStruct/queue/qdata"
	"fmt"
	"os"
)

func generatePrintBinary(n int) {
	if n <= 0 {
		fmt.Println("参数错误")
		os.Exit(2)
	}
	mq := new(qdata.MyQueue)
	mq.Push("1")
	for n > 0 {
		bnum := mq.Pop().(string)
		fmt.Println(bnum)
		mq.Push(bnum + "0")
		mq.Push(bnum + "1")
		n -= 1
	}
}

func main() {
	generatePrintBinary(10)
}
