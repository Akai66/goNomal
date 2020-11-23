package main

import (
    "os"
    "fmt"
)

func main(){
    for i,v := range os.Args[:] {
        fmt.Println(i,":",v)
    }
}
