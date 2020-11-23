package main

import (
    "fmt"
)

func main(){
    fmt.Println(comma("123456"))
}

func comma(s string) string {
    l := len(s)
    if l < 3 {
        return s
    }
    return comma(s[:l-3]) + "," + s[l-3:]
}
