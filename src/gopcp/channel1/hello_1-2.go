package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//总结：导入包时，会优先执行被导入包的init方法，变量（非局部变量）首字母大写才是公开的，才能被包外文件引用，否则只能被包内文件引用
	inputReader := bufio.NewReader(os.Stdin) //带缓冲读取器，从标准输入进行读取
	fmt.Println("Please input your name:")
	input, err := inputReader.ReadString('\n') //以换行符结束一次读取，这个参数必须是字符类型，不能是字符串类型，所以是单引号
	if err != nil {
		fmt.Printf("Found an error:%s\n", err)
	} else {
		input = input[:len(input)-1] //切片操作的第一个参数的含义是从第几个元素开始切(索引从0开始计算，包括该索引值对应的元素，不填默认为0)，第二个参数的含义是切片结束的元素索引(不包含该索引值对应的元素，最大为容量值减一)
		fmt.Printf("hello %s!\n", input)
	}
}
