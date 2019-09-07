package main

import (
	"algorithm/dataStruct/stack/data"
	"fmt"
)

func main() {
	matchBrackets("srtta(jfja{})[")
}

//利用栈实现括号匹配
func matchBrackets(str string) bool {
	mst := new(data.MyStack)
	for _, v := range str {
		if v == '{' || v == '[' || v == '(' {
			mst.Push(v)
		} else if v == '}' || v == ']' || v == ')' {
			//判断栈是否为空
			if mst.Empty() {
				fmt.Println("右括号过多")
				return false
			}
			top := mst.Top().(int32)
			if (top == '{' && v == '}') || (top == '[' && v == ']') || (top == '(' && v == ')') {
				mst.Pop()
			} else {
				fmt.Println("匹配错误")
				return false
			}
		}
	}
	if !mst.Empty() {
		fmt.Println("左括号过多")
		return false
	}
	fmt.Println("匹配成功")
	return true
}
