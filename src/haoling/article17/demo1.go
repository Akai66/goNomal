package main

import "fmt"

func main() {
	sign := make(chan struct{}, 10)
	num := 10
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Println(i)
			sign <- struct{}{}
		}(i)
	}

	for i := 0; i < num; i++ {
		<-sign
	}
}
