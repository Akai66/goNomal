package main

import (
	"fmt"
	. "haoling/article5"
)

var Test int64 = 3 //如果以.的形式导入，有重名变量时，本包的变量会覆盖导入包的变量

func main() {
	fmt.Println(Test)

	var m map[string]int //对于一个值为nil的map，除了添加键-元素对，我们在一个值为nil的字典上做任何操作都不会引起错误
	//m["a"] = 1   //会报panic
	//m["b"]++
	fmt.Printf("%T", m)
}
