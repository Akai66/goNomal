package main

import (
    "os"
    "fmt"
    "strings"
)

func main(){
    fmt.Println(strings.Join(os.Args[:]," "))
    fmt.Println(os.Args[:])
}
