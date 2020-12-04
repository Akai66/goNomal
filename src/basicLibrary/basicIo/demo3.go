package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Copy和CopyN函数
	n, err := io.Copy(os.Stdout, strings.NewReader("Go语言中文网\n"))
	fmt.Printf("copy %d个字节,err:%v\n", n, err)

	//可以直接从输入到输出
	n, err = io.Copy(os.Stdout, os.Stdin)
	fmt.Printf("copy %d个字节,err:%v\n", n, err)
}
