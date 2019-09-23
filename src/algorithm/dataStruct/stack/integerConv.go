package main

import (
	"algorithm/dataStruct/stack/data"
	"fmt"
)

//题目：(整数部分)将十进制数n(正数，负数，0)，转换为R进制数（2<=R<=16,R!=10)
func integerConv(n int, R int) {
	if R >= 2 && R <= 16 && R != 10 {
		ms := new(data.MyStack)
		//是否为负数
		isNegative := false
		if n == 0 {
			fmt.Print(0)
			return
		}
		if n < 0 {
			isNegative = true
			n = -n
		}
		for n != 0 {
			odd := n % R
			switch odd {
			case 10:
				odd = 'A'
			case 11:
				odd = 'B'
			case 12:
				odd = 'C'
			case 13:
				odd = 'D'
			case 14:
				odd = 'E'
			case 15:
				odd = 'F'
			default:
				odd += 48 //字符编码，字符0的编码为48，1的编码为49，以此类推
			}
			ms.Push(odd)
			n /= R
		}
		if isNegative {
			fmt.Print("-")
		}
		for !ms.Empty() {
			fmt.Printf("%c", ms.Pop())
		}
		return
	}
	fmt.Println("参数错误")
}

func main() {
	integerConv(1000, 16)
}
