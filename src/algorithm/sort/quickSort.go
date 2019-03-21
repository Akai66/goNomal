package main

import "fmt"

func main() {
	arr := []int{4, 2, 1, 8, 9, 7}
	arr = quickSort(arr)
	fmt.Println(arr)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := arr[0]
	var leftArr, rightArr []int
	for i := 1; i < len(arr); i++ {
		if arr[i] <= mid {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}
	finalArr := append(quickSort(leftArr), mid)
	return append(finalArr, quickSort(rightArr)...)
}
