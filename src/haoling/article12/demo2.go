package main

import (
	"fmt"
	"os"
)

//高阶函数
//1.接受其它的函数作为参数传入
//2.把其它函数作为结果返回
//至少满足上面其中一个条件的称为高阶函数

type operate func(x, y int) int //函数的签名只和参数列表与结果列表的类型和顺序有关系，和参数名称没关系

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, fmt.Errorf("operate invalid")
	}
	return op(x, y), nil
}

func genOperate(opstr string) (op operate) {
	switch opstr {
	case "+":
		op = func(x, y int) int {
			return x + y
		}
	case "-":
		op = func(x, y int) int {
			return x - y
		}
	case "*":
		op = func(x, y int) int {
			return x * y
		}
	case "/":
		op = func(x, y int) int {
			if y == 0 {
				return 0
			}
			return x / y
		}
	default:

	}
	return
}

func main() {
	//1.接受其它的函数作为参数传入
	op := func(x, y int) int {
		return x + y
	}
	res, err := calculate(1, 2, op)
	if err != nil {
		fmt.Fprintf(os.Stderr, "calculate error:%v\n", err)
		return
	}
	fmt.Println(res)

	//2.把其它函数作为结果返回
	x, y := 4, 2
	calop := genOperate("+")
	fmt.Printf("x+y=%d\n", calop(x, y))
	calop = genOperate("-")
	fmt.Printf("x-y=%d\n", calop(x, y))
	calop = genOperate("*")
	fmt.Printf("x*y=%d\n", calop(x, y))
	calop = genOperate("/")
	fmt.Printf("x/y=%d\n", calop(x, y))
}
