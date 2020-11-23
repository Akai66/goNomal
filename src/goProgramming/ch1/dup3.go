//通过ioutil.ReadFile一次性读取文件全部内容
package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)

func main(){
    counts := make(map[string]int)   //make返回的是引用
    files := os.Args[1:]
    for _,file := range files {
        data,err := ioutil.ReadFile(file) //data类型为byte[]
        if err != nil {
            fmt.Fprintf(os.Stderr,"dup3:%v\n",err)
            continue
        }
        for _,line := range strings.Split(string(data),"\n") {
            if len(line) > 1{
                counts[line]++
            }
        }
    }
    for line,n := range counts {
        fmt.Printf("%d\t%s\n",n,line)
    }
}
