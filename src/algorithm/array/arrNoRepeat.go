package main

import "fmt"

func main() {
	arr := []int{1, 3, 1, 4, 5, 3, 6, 8, 10, 4}
	fmt.Println(firstNoRepeat(arr))
}

func firstNoRepeat(arr []int) int {
	result := -1
	hmap := make(map[int]int, len(arr))
	for _, v := range arr {
		hmap[v]++
	}
	for _, v := range arr {
		if hmap[v] == 1 {
			result = v
			break
		}
	}
	return result
}
