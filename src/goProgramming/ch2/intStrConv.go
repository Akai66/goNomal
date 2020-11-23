package main

import (
    "fmt"
    "strconv"
    "os"
)

func main() {
    //整型转字符串
    fmt.Printf("%T\t%[1]s\n",strconv.Itoa(123))
    s := fmt.Sprintf("%d",123)
    fmt.Printf("%T\t%[1]s\n",s)
    //字符串转整
    i,err := strconv.Atoi("123")
    if err != nil {
        fmt.Fprintf(os.Stderr,"Atoi:%s\n",err)
    }
    fmt.Printf("%T\t%[1]d\n",i)
}
