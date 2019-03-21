package main

import "fmt"

func main() {
	arr := []int{1, 4, 6, 9, 12}
	fmt.Println(towSum(arr, 10))
}

func towSum(arr []int, target int) []int {
	//双指针思想，一个指向最小元素，一个指向最大元素
	i, j := 0, len(arr)-1
	for i < j {
		sum := arr[i] + arr[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
