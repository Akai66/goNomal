package main

import (
	"algorithm/dataStruct/stack/data"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := "6 4 + 2.5 / 5 2 * +"
	fmt.Println(calSuffixExp(str))
}

//利用栈计算后缀表达式的值
func calSuffixExp(expStr string) float64 {
	opts := map[string]int{"+": 1, "-": 1, "*": 1, "/": 1}
	expElements := strings.Split(expStr, " ")
	mst := new(data.MyStack)
	for _, ele := range expElements {
		if _, ok := opts[ele]; ok {
			//如果是操作符则从栈中弹出两个元素，进行计算后，将结果入栈,注意：在弹出时被减(加、乘、除)数先被弹出
			num1 := mst.Top()
			mst.Pop()
			num2 := mst.Top()
			mst.Pop()
			result := 0.0
			switch ele {
			case "+":
				result = num2 + num1
			case "-":
				result = num2 - num1
			case "*":
				result = num2 * num1
			case "/":
				result = num2 / num1
			}
			//将结果入栈
			mst.Push(result)
		} else {
			//如果是操作数则直接入栈
			v, err := strconv.ParseFloat(ele, 64)
			if err != nil {
				fmt.Println("后缀表达式格式错误，操作数为float类型，操作符集合：+,-,*,/,元素间用空格分隔")
				os.Exit(1)
			}
			mst.Push(v)
		}
	}
	return mst.Top()
}
