//使用io.Copy方法直接将body内容输出到标准输出
//兼容http://前缀
//输出http响应码

package main

import (
    "fmt"
    "os"
    "io"
    "net/http"
    "strings"
)

func main(){
    for _,url := range os.Args[1:] {
        if !strings.HasPrefix(url,"http://") {
            url = "http://" + url
        }
        srep,err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr,"fetch:%v\n",err)
            os.Exit(1)
        }
        defer  srep.Body.Close()
        n,err := io.Copy(os.Stdout,srep.Body)
        if err != nil {
            fmt.Fprintf(os.Stderr,"fetch:copying %s:%v\n",url,err)
            os.Exit(1)
        }
        fmt.Printf("\ncontent length:%v,res status:%v\n",n,srep.Status)
    }
}
