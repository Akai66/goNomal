//根据命令行参数传入的标题，从离线索引中获取对应的漫画url
package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "goProgramming/ch4/xkcd"
)

func main(){
    var title string
    if len(os.Args)>1 {
        title = os.Args[1]
    }
    f,err := os.Open("comicIndex.json")
    if err != nil {
        fmt.Fprintf(os.Stderr,"file open failed:%v\n",err)
        os.Exit(1)
    }
    defer f.Close()
    data,err := ioutil.ReadAll(f)
    if err != nil {
        fmt.Fprintf(os.Stderr,"read file to data failed:%v\n",err)
        os.Exit(1)
    }
    comicsMap := make(map[string]xkcd.Comic)
    if err := json.Unmarshal(data,&comicsMap);err != nil {
        fmt.Fprintf(os.Stderr,"json unmarshal failed:%v\n",err)
        os.Exit(1)
    }
    if comic,ok := comicsMap[title];ok {
        fmt.Printf("%s ===> %s\n",title,comic.Img)
    }else{
        fmt.Printf("%s ===> %s\n",title,"not found!")
    }
}

