package main

import (
    "fmt"
    "bytes"
)

func main(){
    fmt.Println(intsToString([]int{1,2,3}))
}

func intsToString(value []int) string {
    var bf bytes.Buffer
    bf.WriteByte('[')
    for i,v := range value {
        if i > 0 {
            bf.WriteString(", ")
        }
        fmt.Fprintf(&bf,"%d",v)
    }
    bf.WriteByte(']')
    return bf.String()
}
