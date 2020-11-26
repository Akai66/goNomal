package main

import "fmt"

type Printer func(string) (int, error)

func printToStd(content string) (int, error) {
	return fmt.Println(content)
}

func main() {
	//函数可以作为值赋值给变量
	var p Printer
	p = printToStd
	p("hello world")
}
