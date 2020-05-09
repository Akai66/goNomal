package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//对非http://开头的，默认添加该前缀
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Println(url)
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		//打印http响应码
		fmt.Printf("http status code : %s\n", res.Status)

		//先读取到缓冲区
		//data, err := ioutil.ReadAll(res.Body)
		//res.Body.Close()
		//if err != nil {
		//	fmt.Fprintf(os.Stderr, "fetch:reading %s:%v\n", url, err)
		//	os.Exit(1)
		//}
		//fmt.Printf("%s", data)

		//利用io.Copy将响应内容输入流直接copy到标准输出
		io.Copy(os.Stdout, res.Body)
		res.Body.Close()
	}
}
