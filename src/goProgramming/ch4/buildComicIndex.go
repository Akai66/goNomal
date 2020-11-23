package main

import (
	"fmt"
	"goProgramming/ch4/xkcd"
	"os"
	"strconv"
    "encoding/json"
)

func main() {
	num := 10
	if len(os.Args) > 1 {
		var err error
		num, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "params err:%v\n", err)
		}
	}
	comics, err := xkcd.GetComics(num)
	if err != nil {
		fmt.Fprintf(os.Stderr, "get comics err:%v\n", err)
	}
    //将切片转为map，然后进行json编码
    comicsMap := make(map[string]xkcd.Comic)
    for _,comic := range comics {
        comicsMap[comic.Title] = comic
    }
    data,_ := json.MarshalIndent(comicsMap,""," ")
	//将结果写入本地离线文件
    f,err := os.OpenFile("comicIndex.json",os.O_RDWR|os.O_TRUNC|os.O_CREATE,0755) //OpenFile打开文件有个小坑，如果需要覆盖写入，需要添加os.O_TRUNC参数
    if err != nil {
        fmt.Fprintf(os.Stderr,"open file failed:%v\n",err)
        os.Exit(1)
    }
    defer f.Close()
    _,err = f.Write(data)
    if err != nil {
        fmt.Fprintf(os.Stderr,"file write failed:%v\n",err)
        os.Exit(1)
    }
}
