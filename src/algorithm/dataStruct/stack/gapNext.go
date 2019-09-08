package main

import (
	"algorithm/dataStruct/stack/data"
	"fmt"
)

func main() {
	arr := []int{15, 1, 3, 9, 10, 7, 6, 5, 12, 2}
	fmt.Println(gapNextBigger(arr))
}

//数组中元素与下一个比它大的元素之间的距离
func gapNextBigger(arr []int) (dist []int) {
	dist = make([]int, len(arr))
	ms := new(data.MyStack)
	for curIndex, curValue := range arr {
		for !ms.Empty() && curValue > arr[ms.Top().(int)] {
			preIndex := ms.Pop().(int)
			dist[preIndex] = curIndex - preIndex
		}
		ms.Push(curIndex)
	}
	return dist
}
