package main

import "fmt"

//给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
//
//返回 s 所有可能的分割方案。

//示例:
//
//输入: "aab"
//输出:
//[
//  ["aa","b"],
//  ["a","a","b"]
//]

func splitPalindrome(str string) [][]string {
	res := [][]string{}
	if len(str) <= 0 {
		return res
	}
	size := len(str)
	for i := 1; i <= size; i++ {
		substr := str[0:i] //截取[0,i)
		if isPalindromeSimple(substr) {
			temp := []string{substr}
			part := splitPalindrome(str[i:])
			res = merge(res, temp, part)
		}
	}
	return res
}

func isPalindromeSimple(str string) bool {
	i := 0
	j := len(str) - 1
	for i < j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

//值传递：改变形参的值，并不会改变实参的值。（例如：数组、结构体）
//地址传递(引用传递)：改变形参的值，会影响到实参的值。（例如：切片、map、指针）（切片变量名本身就是一个地址） 特别注意：如果是append()添加数据时不会影响实参
func merge(res [][]string, temp []string, part [][]string) [][]string {
	if len(part) <= 0 {
		res = append(res, temp)
	} else {
		for _, arr := range part {
			data := temp
			for _, value := range arr {
				data = append(data, value)
			}
			res = append(res, data)
		}
	}
	return res
}

func main() {
	str := "eaabbaea"
	fmt.Println(splitPalindrome(str))
}
