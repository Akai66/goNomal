 package main

import (
    "fmt"
    "unsafe"
)

func main(){
    var i int32 = 20
    fmt.Printf("变量i的类型为%T,占用的字节数是%d个字节\n",i,unsafe.Sizeof(i))
    var ret = false
    fmt.Printf("ret的类型为%T,占用字节为%d个字节\n",ret,unsafe.Sizeof(ret))
    srcData := []int{1,2,3,4,5}
    data := make([]int,10)
    copy(data,srcData)
    fmt.Println(data)
    copy(data[5:],srcData[2:4])
    fmt.Println(data)
}
