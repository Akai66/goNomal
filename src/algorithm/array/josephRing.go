package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(josephRing(arr, 3))
}

func josephRing(arr []int, n int) int {
	i := 1
	for ; len(arr) > 1; i++ {
		v := arr[0]
		arr = arr[1:]
		if i%n != 0 {
			arr = append(arr, v)
		}
	}
	return arr[0]
}
