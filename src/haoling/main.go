package main

import (
	"fmt"
	. "haoling/article5"
)

var Test int64 = 3   //如果以.的形式导入，有重名变量时，本包的变量会覆盖导入包的变量

func main() {
	fmt.Println(Test)
}
