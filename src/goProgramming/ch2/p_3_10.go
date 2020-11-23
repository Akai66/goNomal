//使用非递归方式,bytes.Buffer实现comma函数
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1234567"))
}

func comma(s string) string {
	var bf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			bf.WriteByte(',')
		}
		bf.WriteByte(s[i])
	}
	return bf.String()
}
