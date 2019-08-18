package main

import "fmt"

func main() {
	arr := []int{6, 2, 5, 1, 9, 10, 1, 3, 8, 4}
	heapSort(arr)
	fmt.Println(arr)
}

//堆排序
func heapSort(arr []int) {
	arrLen := len(arr)
	buildMaxHeap(arr, arrLen)
	for i := arrLen - 1; i >= 0; i-- {
		swap(arr, 0, i)
		arrLen--
		heapify(arr, 0, arrLen)
	}
}

//构造大顶堆
func buildMaxHeap(arr []int, arrLen int) {
	for i := arrLen/2 - 1; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

//以i为根节点，进行调整，让子树满足大顶堆规则
func heapify(arr []int, i, arrLen int) {
	lcIndex := i*2 + 1
	rcIndex := i*2 + 2
	maxIndex := i
	if lcIndex < arrLen && arr[lcIndex] > arr[maxIndex] {
		maxIndex = lcIndex
	}
	if rcIndex < arrLen && arr[rcIndex] > arr[maxIndex] {
		maxIndex = rcIndex
	}
	if maxIndex != i {
		swap(arr, i, maxIndex)
		heapify(arr, maxIndex, arrLen)
	}
}

//交换
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
