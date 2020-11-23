//编写一个函数，在原地完成消除[]string中相邻重复的字符串的操作
package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c", "b", "d", "d", "d", "e", "e"}
	fmt.Println(s)
	res := delDupStr(s)
	fmt.Println(res)
}

func delDupStr(s []string) []string {
	oriLen := len(s)
	for i := 0; i < oriLen; i++ {
		index := i + 1
		newLen := len(s)
		if index >= newLen {
			break
		}
		if s[i] == s[index] {
			copy(s[i:], s[index:])
			s = s[:newLen-1]
			i--
		}
	}
	return s
}
