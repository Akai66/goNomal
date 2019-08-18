package main

import "fmt"

func main() {
	arr := [...]int{10, 2, 9, 3, 7, 1}
	bubbleSort(&arr)
	fmt.Println(arr)
}

func bubbleSort(arr *[6]int) {
	length := len(*arr)
	for i := 0; i < length-1; i++ {
		flag := true
		for j := 0; j < length-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}
