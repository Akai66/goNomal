//defer,panic,recover 内置函数的使用
package main

import (
	"fmt"
)

func main() {
	printNum()
	fmt.Println()
	v, err := getValue(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

func printNum() {
	for i := 1; i < 6; i++ {
		//延迟函数，栈，先进后出
		defer func(n int) {
			fmt.Print(n)
		}(i)
	}
	fmt.Println("此行优先执行")
	return
}

func getValue(index int) (n int, err error) {
	defer func() {
		//recover内置函数会拦截运行时恐慌，使当前程序从恐慌状态中恢复并重新获得流程控制权，恐慌不会再蔓延到上层调用函数，除非defer函数在拦截恐慌后又重新panic恐慌
		if p := recover(); p != nil {
			if se, ok := p.(error); ok {
				err = se
			} else {
				panic(p)
			}

		}
	}()
	arr := [3]int{1, 2, 3}
	n = arr[index]
	return
}
