package main

import (
    "fmt"
    "log"
    "net/http"
)

type dollars float32

func (d dollars) String() string {return fmt.Sprintf("$%.2f",d)}

type database map[string]dollars

func (db database) ServeHTTP (w http.ResponseWriter,req *http.Request) {
    switch req.URL.Path {
        case "/list":
            for item,price := range db {
                fmt.Fprintf(w,"%s:%s\n",item,price)
            }
        case "/price":
            item := req.URL.Query().Get("item")
            if price,ok := db[item];ok {
                fmt.Fprintf(w,"%s\n",price)
            }else{
                msg := fmt.Sprintf("no such item:%q\n",item)
                http.Error(w,msg,http.StatusNotFound)
            }
        default:
            msg := fmt.Sprintf("no such page:%s\n",req.URL)
            http.Error(w,msg,http.StatusNotFound)
    }
}

func main(){
    db := database{"shoes":50,"mp3":100,"iphone":1000}
    log.Fatal(http.ListenAndServe("localhost:8000",db))
}
