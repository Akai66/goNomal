package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main(){
    file1,file2 := os.Args[1],os.Args[2]
    f1,_ := os.Open(file1)
    f2,_ := os.Open(file2)
    defer f1.Close()
    defer f2.Close()
    input1 := bufio.NewScanner(f1)
    input2 := bufio.NewScanner(f2)
    for input1.Scan() {
        input2.Scan()
        freq,_ := strconv.Atoi(input2.Text())
        freq *= 1000
        fmt.Printf("update adver_ssp_strategy set total_freq_limit = '{\"num\":%d,\"type\":\"0\",\"cycleType\":\"0\",\"startdate\":\"2020-11-01\",\"enddate\":\"2020-11-30\"}' where id = %s;\n",freq,input1.Text())
    }
}
