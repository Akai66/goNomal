package main

import (
	"fmt"
	"strings"
)

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写

func isPalindrome(str string) bool {
	//先转换为小写
	str = strings.ToLower(str)
	i := 0
	j := len(str) - 1
	for i < j {
		if isLetterOrDigit(str[i]) && isLetterOrDigit(str[j]) {
			if str[i] == str[j] {
				i++
				j--
			} else {
				return false
			}
		} else {
			if !isLetterOrDigit(str[i]) {
				i++
			}
			if !isLetterOrDigit(str[j]) {
				j--
			}
		}
	}
	return true
}

func isLetterOrDigit(ch uint8) bool {
	if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}

func main() {
	str := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
	str = "Aba"
	fmt.Println(isPalindrome(str))
	str = "race a car"
	fmt.Println(isPalindrome(str))
}
