package main

import (
	"fmt"
	"sort"
)

//判断arr2是否为arr1的子集
/**
先排序，然后通过归并的方式判断
*/
func isSubArray1(arr1, arr2 []int) bool {
	sort.Ints(arr1)
	sort.Ints(arr2)
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	i, j := 0, 0
	for i < arr1Len && j < arr2Len {
		if arr1[i] < arr2[j] {
			i++
		} else if arr1[i] == arr2[j] {
			i++
			j++
		} else {
			return false
		}
	}
	if j < arr2Len {
		return false
	}
	return true
}

/**
利用hash表，进行判断
*/
func isSubArray2(arr1, arr2 []int) bool {
	imap := map[int]int{}
	for _, v := range arr1 {
		imap[v] += 1
	}
	for _, v := range arr2 {
		if imap[v] > 0 {
			imap[v] -= 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{11, 9, 4, 10, 3, 16}
	arr2 := []int{4, 11, 10, 9, 3}
	fmt.Println(isSubArray1(arr1, arr2))
	fmt.Println(isSubArray2(arr1, arr2))
}
