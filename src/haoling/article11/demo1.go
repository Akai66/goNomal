package main

import (
	"fmt"
	"time"
)

func main() {
	//通道select用法
	ch := make(chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(ch)
	})
	select {
	case _, ok := <-ch:
		if !ok {
			fmt.Println("the candidate case is closed")
			break
		}
		fmt.Println("the candidate case is selected")
	}
}
