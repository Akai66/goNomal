package main

import "fmt"

var chs = [3]chan int{
	nil,
	make(chan int),
	nil,
}

var nums = []int{1, 2, 3}

func main() {
	for i := 0; i < 2; i++ {
		select {
		case getChan(0) <- getNum(0):
			fmt.Println("the first candidate case is selected")
		case getChan(1) <- getNum(1):
			fmt.Println("the second candidate case is selected")
		case getChan(2) <- getNum(2):
			fmt.Println("the third candidate case is selected")
		default:
			fmt.Println("no candidate case is selected")
		}
	}
}

func getNum(i int) int {
	fmt.Printf("Number:%d\n", i)
	return nums[i]
}

func getChan(i int) chan int {
	fmt.Printf("Chan:%d\n", i)
	return chs[i]
}
