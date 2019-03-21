package main

import (
    "fmt"
    "time"
    "strconv"
)

func joinStr() {
    var l string
    for i:=0;i<100000;i++ {
        l += strconv.Itoa(i)
    }
}

func main() {
    start := time.Now().Unix()
    joinStr()
    end := time.Now().Unix()
    fmt.Printf("joinStr方法执行时间:%d秒",end-start)
}
