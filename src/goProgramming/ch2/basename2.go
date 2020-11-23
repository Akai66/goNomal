package main

import (
    "fmt"
    "strings"
    "os"
)

func main(){
    path := os.Args[1]
    fmt.Println(basename(path))
    s := "abc我"
    fmt.Println(string([]rune(s)[3]))
}

func basename(s string) string{
    slash := strings.LastIndex(s,"/")
    s = s[slash+1:]
    dot := strings.LastIndex(s,".")
    if dot > 0 {
        s = s[:dot]
    }
    return s
}
