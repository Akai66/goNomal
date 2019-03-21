package main

import "fmt"

func main() {
	arr := []int{10, 3, 8, 2, 19, 27, 23, 4}
	fmt.Println(secondMin(arr))
}

func secondMin(arr []int) int {
	if len(arr) < 2 {
		return -1
	}
	firstMin, secondMin := arr[0], arr[0]
	for _, v := range arr {
		if v < firstMin {
			secondMin = firstMin
			firstMin = v
		} else if v < secondMin && v != firstMin {
			secondMin = v
		}
	}
	if firstMin == secondMin {
		return -1
	}
	return secondMin
}
