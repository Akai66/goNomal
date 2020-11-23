//编写一个函数，原地将一个UTF8编码的[]byte类型的slice中相邻的空格替换成一个空格

package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []byte{'a', ' ', ' ', 'b', 'c', ' ', 'd'}
	fmt.Println(string(s))
	s = dupSpace(s)
	fmt.Println(string(s))
}

func dupSpace(s []byte) []byte {
	oriLen := len(s)
	for i := 0; i < oriLen; i++ {
		index := i + 1
		newLen := len(s)
		if index >= newLen {
			break
		}
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[index])) {
			copy(s[i:], s[index:])
			s = s[:newLen-1]
			i--
		}
	}
	return s
}
