//通过bufio的Scanner一行行读取文件内容
package main

import (
    "os"
    "fmt"
    "bufio"
)

func main(){
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin,counts)
    }else{
        for _,file := range files {
            fd,err := os.Open(file)
            if err != nil {
                fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
                continue
            }
            countLines(fd,counts)
            fd.Close()
        }
    }

    //输出结果
    for line,n := range counts {
        fmt.Printf("%d\t%s\n",n,line)
    }
}

func countLines(f *os.File,counts map[string]int){
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}
