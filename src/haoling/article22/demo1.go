package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("enter func main")
	defer func() {
		fmt.Println("enter func defer")
		if p := recover(); p != nil {
			panic(fmt.Errorf("fatal error:%s", p))
		}
		fmt.Println("exit func defer")
	}()
	panic(errors.New("something error"))
	fmt.Println("exit func main")
}
