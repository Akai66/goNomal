package main

import "fmt"

func main() {
	arr := [...]int{2, 7, 1, 3, 5, 4}
	selectSort(&arr)
	fmt.Println(arr)
}

func selectSort(arr *[6]int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if (*arr)[j] < (*arr)[i] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[i]
				(*arr)[i] = temp
			}
		}
	}
}
