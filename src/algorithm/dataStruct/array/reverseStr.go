package main

import "fmt"

func main() {
	fmt.Println(reverse("我喜欢go"))
}

func reverse(str string) string {
	if len(str) < 1 {
		return str
	}
	//转换为切片，因为如果存在中文，byte数组会出现乱码
	strRune := []rune(str)
	strLen := len(strRune)
	mid := strLen / 2
	for i := 0; i <= mid; i++ {
		temp := strRune[i] //temp必须使用预判赋值
		strRune[i] = strRune[strLen-1-i]
		strRune[strLen-1-i] = temp
	}
	return string(strRune)
}
