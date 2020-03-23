package main

import "fmt"

/*
题目描述:找出数组中的重复数字
在一个长度为n的数组里的所有数字都在0~n-1的范围内，数组中某些数字是重复的，找出数组中任意一个重复的数字
*/

func main() {
	arr := []int{1, 3, 2, 4, 5, 3, 2, 7}
	ret := duplicate(arr)
	fmt.Println(ret)
}

func duplicate(arr []int) int {
	length := len(arr)
	if length <= 0 {
		fmt.Print("参数错误:数组为空")
		return -1
	}
	for i := 0; i < length; i++ {
		if arr[i] < 0 || arr[i] > length-1 {
			fmt.Print("参数错误:数组元素不符合要求")
			return -1
		}
	}
	for i := 0; i < length; i++ {
		for arr[i] != i {
			if arr[i] == arr[arr[i]] {
				return arr[i]
			} else {
				//将数字arr[i]交换到对应的数组位置上去，如：数字2交换到arr[2]位置
				arr[i], arr[arr[i]] = arr[arr[i]], arr[i]
			}
		}
	}
	return -1
}
