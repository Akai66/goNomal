package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Printf("Map: %v\n", m)
	info = fmt.Sprintf("OS:%s,Arch:%s", runtime.GOOS, runtime.GOARCH)
}

var m = map[int]string{1: "A", 2: "B", 3: "C"}

var info string

func main() {
	//总结：程序执行顺序：全局变量初始化 优先于 init函数 优先于 main函数
	//原因：如果init函数优先级最高，那么全局变量在init函数中被修改后，又会被全局变量初始化覆盖，所以全局变量初始化的优先级高于init函数
	//同一个代码包可以有多个代码包初始化函数，但是go不会保证同一个代码包所有初始化函数的执行顺序，被导入的代码包的初始化函数总是优先于本文件的代码包初始化函数执行
	fmt.Println(info)
}
