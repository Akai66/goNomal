package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(1.1))
}

func sqrt(n float64) float64 {
	if n < 1 { //小于1的数的平方根>原数，二分法无法解决
		return -1.0
	}
	low, high := 0.0, n
	mid := (low + high) / 2
	for math.Abs(mid*mid-n) > 1e-6 {
		if mid*mid > n {
			high = mid
		} else if mid*mid < n {
			low = mid
		} else {
			return mid
		}
		mid = (low + high) / 2
	}
	return mid
}
