package main

import "fmt"

func main() {
	var num int
	for i := 10; i > 0; i-- {
		num = peach(i)
		fmt.Printf("第%d天有%d个桃子\n", i, num)
	}
}

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("input error!!")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
