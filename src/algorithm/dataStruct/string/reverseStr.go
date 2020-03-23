package main

import "fmt"

func reverse(str string) string {
	if len(str) < 1 {
		return str
	}
	//转换为切片，因为如果存在中文，byte数组会出现乱码
	strRune := []rune(str)
	i, j := 0, len(strRune)-1
	for i < j {
		strRune[i], strRune[j] = strRune[j], strRune[i]
		i++
		j--
	}
	return string(strRune)
}

func main() {
	fmt.Println(reverse("loveandpeace"))
}
