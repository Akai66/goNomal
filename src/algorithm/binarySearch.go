package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(arr, 2))
}

func binarySearch(arr []int, target int) int {
	min, max := 0, len(arr)-1
	for min <= max {
		mid := (min + max) / 2
		if arr[mid] == target {
			return mid + 1
		} else if arr[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return -1
}
