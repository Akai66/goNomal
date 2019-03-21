package main

import "fmt"

func main() {
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = feibo(i + 1)
	}
	fmt.Printf("斐波那契队列的前10个数字为:%v\n", arr)
}

func feibo(n int) int {
	if n < 1 { //边界判断
		fmt.Println("input error")
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	} else {
		return feibo(n-1) + feibo(n-2)
	}
}
