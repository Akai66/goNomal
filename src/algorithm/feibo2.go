package main

import "fmt"

func main() {
	fb := feibo(10)
	fmt.Println(fb)
}

func feibo(n int) []int {
	//切片
	var fb []int = make([]int, n)
	if n == 1 {
		fb[0] = 1
	} else {
		fb[0] = 1
		fb[1] = 1
		if n > 2 {
			for i := 2; i < n; i++ {
				fb[i] = fb[i-1] + fb[i-2]
			}
		}
	}
	return fb
}
