package main

import "fmt"

func main() {
	arr := []int{10, 6, 3, 8, 7, 4, 5}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && current < arr[preIndex] {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}
