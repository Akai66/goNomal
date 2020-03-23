package main

import (
	"fmt"
)

//将一个字符串转换成一个整数，字符串不是一个合法的数值则返回 0，要求不能使用字符串转换整数的库函数。

//指针是否为空指针以及字符串是否为空字符串；
//
//字符串对于正负号的处理；
//
//输入值是否为合法值，即小于等于'9'，大于等于'0'；
//
//int为32位，需要判断是否溢出；
//
//使用错误标志，区分合法值0和非法值0。

func strToInt(str string) (ret int, valid bool) {
	ret, valid = 0, false
	strLen := len(str)
	if strLen < 1 || (strLen == 1 && (str == "+" || str == "-")) {
		//字符串为空,"+","-",均不合法
		return
	}
	var isPositive int = 1
	if str[0] == '-' {
		isPositive = -1
	}
	var i int
	for i = 0; i < strLen; i++ {
		if i == 0 && (str[i] == '+' || str[i] == '-') {
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			ret = 0
			break
		}
		ret = ret*10 + isPositive*int(str[i]-'0')
	}

	if i >= strLen {
		valid = true
	}
	return
}

func main() {
	str := "-1024"
	fmt.Println(strToInt(str))
}
