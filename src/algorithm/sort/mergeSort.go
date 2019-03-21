package main

import "fmt"

func main() {
	arr := []int{2, 8, 1, 6, 3, 5, 7}
	arr = sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftArr := arr[:mid]
	rightArr := arr[mid:]
	return merge(sort(leftArr), sort(rightArr))
}

func merge(leftArr []int, rightArr []int) []int {
	var resultArr []int
	for len(leftArr) > 0 && len(rightArr) > 0 {
		if leftArr[0] > rightArr[0] {
			resultArr = append(resultArr, rightArr[0])
			rightArr = rightArr[1:]
		} else {
			resultArr = append(resultArr, leftArr[0])
			leftArr = leftArr[1:]
		}
	}
	if len(leftArr) > 0 {
		resultArr = append(resultArr, leftArr...)
	}
	if len(rightArr) > 0 {
		resultArr = append(resultArr, rightArr...)
	}
	return resultArr
}
