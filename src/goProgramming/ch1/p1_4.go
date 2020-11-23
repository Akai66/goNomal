//分别统计每个文件内重复的行数
package main

import (
    "os"
    "bufio"
    "fmt"
)

func main(){
    files := os.Args[1:]
    for _,file := range files {
        fd,err := os.Open(file)
        if err != nil {
            fmt.Fprintf(os.Stderr,"p1_4:%v\n",err)
            continue
        }
        countLine(fd,file)
        fd.Close()
    }
}

func countLine(f *os.File,fn string){
    counts := make(map[string]int)
    fmt.Println(fn)
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
    for line,n := range counts {
        fmt.Printf("%d\t%s\n",n,line)
    }
}
