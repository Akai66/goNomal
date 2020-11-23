//修改reverse函数用于原地反转UTF8编码的[]byte,是否可以不用分配额外的内存

package main

import (
	"fmt"
)

func main() {
	s := []byte{'a', 'b', 'c'}
	fmt.Println(string(s))
	reverse(s)
	fmt.Println(string(s))
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
