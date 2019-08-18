package main

import "fmt"

type util struct {
}

func main() {
	var arr [][]int = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {4, 7, 2, 0}, {7, 9, 1, 4}}
	u := new(util)
	u.printArr(arr)
	arr = zhuanzhi(arr)
	fmt.Println()
	u.printArr(arr) //go底层自动转换为(*u).printArr(arr)
}

func (u util) printArr(arr [][]int) {
	for _, a := range arr {
		for _, v := range a {
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func zhuanzhi(arr [][]int) [][]int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr[i]); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	return arr
}
