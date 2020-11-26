package main

import "fmt"

func main() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
		}
	}
	fmt.Println(numbers1)

	numbers2 := [...]int{1, 2, 3, 4, 5, 6} //数组，值类型
	maxIndex := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

	numbers3 := []int{1, 2, 3, 4, 5, 6} //切片，引用类型
	maxIndex = len(numbers3) - 1
	for i, e := range numbers3 {
		if i == maxIndex {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	fmt.Println(numbers3)
}
