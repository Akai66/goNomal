package main

import (
	"os"
	"fmt"
)

func main()  {
	for i,v := range os.Args[1:] {
		fmt.Printf("index:%v,value:%v\n",i,v)
	}
}
