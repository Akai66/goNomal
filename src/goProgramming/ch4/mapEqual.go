//判断两个map是否相等，判断长度是否相等，以及是否包含相同的key和value
package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	y := map[string]int{"b": 2, "d": 4, "a": 1}
	fmt.Println(equal(x, y))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk]; !ok || yv != xv {
			return false
		}
	}
	return true
}
