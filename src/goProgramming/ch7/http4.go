//支持增删改查接口

package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type dollars float64

func (d dollars) String() string {return fmt.Sprintf("$%.2f",d)}

type database map[string]dollars

func (db database) add(w http.ResponseWriter,r *http.Request){
   key,value := r.URL.Query().Get("key"),r.URL.Query().Get("value")
   fv,err := strconv.ParseFloat(value,64)
   if err != nil {
       fmt.Fprintf(w,"params value error,must be float:%q\n",value)
       return
   }
   db[key] = dollars(fv)
   fmt.Fprintf(w,"add success!")
}

func (db database) del(w http.ResponseWriter,r *http.Request){
    key := r.URL.Query().Get("key")
    delete(db,key)
    fmt.Fprintf(w,"del success!")
}

func (db database) update(w http.ResponseWriter,r *http.Request){
    key := r.URL.Query().Get("key")
    if _,ok := db[key];!ok {
        fmt.Fprintf(w,"key not exist:%q\n",key)
        return
    }
    value := r.URL.Query().Get("value")
    fv,err := strconv.ParseFloat(value,64)
    if err != nil {
        fmt.Fprintf(w,"params value error,must be float:%q\n",value)
        return
    }
    db[key] = dollars(fv)
    fmt.Fprintf(w,"update success!")
}

func (db database) get(w http.ResponseWriter,r *http.Request){
    key := r.URL.Query().Get("key")
    if value,ok := db[key];ok {
        fmt.Fprintf(w,"%s\n",value)
    }else{
        msg := fmt.Sprintf("key not exist:%q\n",key)
        http.Error(w,msg,http.StatusNotFound)
    }
}

func main(){
    db := database{"shoes":50,"iphone":800}
    http.HandleFunc("/get",db.get)
    http.HandleFunc("/add",db.add)
    http.HandleFunc("/del",db.del)
    http.HandleFunc("/update",db.update)
    log.Fatal(http.ListenAndServe("localhost:8000",nil))
}
