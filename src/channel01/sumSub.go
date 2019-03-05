package main

import (
	"fmt"
)

func main() {
	sum, sub := getVal(6, 4)
	fmt.Println("sum=", sum, ";sub=", sub)
	sum1, _ := getVal(3, 6)
	fmt.Println("sum=", sum1)
}

func getVal(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}
