package main

import "fmt"

func main() {
	arr1 := []int{3, 4, 7, 8, 12}
	arr2 := []int{1, 2, 5, 6, 9, 10, 13}
	fmt.Println(mergeSortedArr1(arr1, arr2))
	arr3 := make([]int, 12)
	copy(arr3, arr1)
	fmt.Println(mergeSortedArr2(arr3, 5, arr2, 7))
}

//借助第三个切片
func mergeSortedArr1(arr1, arr2 []int) []int {
	i, j := 0, 0
	var mergeArr []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			mergeArr = append(mergeArr, arr1[i])
			i++
		} else {
			mergeArr = append(mergeArr, arr2[j])
			j++
		}
	}
	if i < len(arr1) {
		mergeArr = append(mergeArr, arr1[i:]...)
	}
	if j < len(arr2) {
		mergeArr = append(mergeArr, arr2[j:]...)
	}
	return mergeArr
}

//不借助第三个切片，从后向前填充
func mergeSortedArr2(arr1 []int, m int, arr2 []int, n int) []int {
	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if arr1[i] >= arr2[j] {
			arr1[k] = arr1[i]
			i--
		} else {
			arr1[k] = arr2[j]
			j--
		}
	}
	if j >= 0 {
		for ; j >= 0; j-- {
			arr1[j] = arr2[j]
		}
	}
	return arr1
}
