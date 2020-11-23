//编写一个rotate函数，通过一次循环完成旋转
package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(rotate(s, 2))
	fmt.Println(rotate(s, 1))
	fmt.Println(rotate(s, 0))
}

func rotate(s []int, r int) []int {
	lens := len(s)
	res := make([]int, lens)
	for i := 0; i < lens; i++ {
		index := i + r
		if index >= lens {
			index = index - lens
		}
		res[i] = s[index]
	}
	return res
}
