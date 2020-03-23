package main

import "fmt"

/*
题目描述：不修改数组找出重复的数字
在一个长度为n+1的数组里，所有数字都在1~n范围内，所以数组中至少有一个数字是重复的，请找出数组中任意一个重复的数字，但不能修改输入的数组
*/
func main() {
	arr := []int{2, 3, 5, 4, 3, 2, 6, 7}
	fmt.Println(duplicate2(arr))
}

//时间复杂度为O(n*logn)
func duplicate2(arr []int) int {
	length := len(arr)
	if length <= 1 {
		fmt.Println("参数错误:数组元素个数必须大于1")
		return -1
	}
	for i := 0; i < length; i++ {
		if arr[i] < 1 || arr[i] > length-1 {
			fmt.Println("参数错误:数组元素不符合要求")
			return -1
		}
	}
	start, end := 1, length-1
	for start <= end {
		mid := (end + start) / 2
		count := countRange(arr, start, mid)
		if start == end {
			if count > 1 {
				return start
			} else {
				break
			}
		}
		if count > mid-start+1 {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

func countRange(arr []int, start int, end int) int {
	count := 0
	length := len(arr)
	if length <= 1 {
		fmt.Println("参数错误:数组元素个数必须大于1")
		return count
	}
	for i := 0; i < length; i++ {
		if arr[i] >= start && arr[i] <= end {
			count++
		}
	}
	return count
}
