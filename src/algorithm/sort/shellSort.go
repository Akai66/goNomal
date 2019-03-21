package main

import "fmt"

func main() {
	arr := []int{7, 3, 1, 10, 19, 30, 9}
	shellSort(arr)
	fmt.Println(arr)
}

func shellSort(arr []int) {
	length := len(arr)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			current := arr[i]
			preIndex := i - gap
			for preIndex >= 0 && current < arr[preIndex] {
				arr[preIndex+gap] = arr[preIndex]
				preIndex -= gap
			}
			arr[preIndex+gap] = current
		}
	}
}
