//函数类型，go语言中函数类型是一等类型，既可以作为其它函数的参数，也可以作为其结果
package main

import (
	"errors"
	"fmt"
)

//定义函数类型binaryOperation
type binaryOperation func(operand1 int, operand2 int) (result int, err error)

//声明函数类型变量，并赋值
var divide binaryOperation = func(operand1 int, operand2 int) (result int, err error) {
	if operand2 == 0 {
		err = errors.New("division by zero")
		return
	}
	result = operand1 / operand2
	return
}

//参数名可以和定义函数类型时的不同，但是参数类型必须相同
var add binaryOperation = func(op1 int, op2 int) (result int, err error) {
	result = op1 + op2
	return
}

var sub = func(op1 int, op2 int) (ret int, err error) {
	ret = op1 - op2
	return
}

//函数闭包
func operate(op1 int, op2 int, bop binaryOperation) (result int, err error) {
	if bop == nil {
		err = errors.New("invalid binary operation function")
		return
	}
	return bop(op1, op2)
}

func main() {
	result, err := operate(4, 0, divide)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result, _ = operate(4, 2, add)
	fmt.Println(result)
	result, _ = operate(3, 6, sub)
	fmt.Println(result)
}
