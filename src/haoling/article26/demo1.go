package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var protecting uint

func init() {
	flag.UintVar(&protecting, "protecting", 1, "It indicates whether to use a mutex to protect data writing")
}

func main() {
	flag.Parse()
	var buffer bytes.Buffer

	const (
		max1 = 5  //代表启用goroutine的数量
		max2 = 10 //代表每个goroutine需要写入的数据块的数量
		max3 = 10 //代表每个数据块中需要写入的重复数字
	)

	var mu sync.Mutex

	sign := make(chan struct{}, max1)

	for i := 1; i <= max1; i++ {
		go func(id int, writer io.Writer) {
			defer func() {
				sign <- struct{}{}
			}()
			for j := 1; j <= max2; j++ {
				header := fmt.Sprintf("\n[id:%d,iteration:%d]", id, j)
				data := fmt.Sprintf(" %d", id*j)
				if protecting > 0 {
					mu.Lock()
				}
				_, err := buffer.Write([]byte(header))
				if err != nil {
					log.Printf("error:%s [%d]", err, id)
				}
				for k := 0; k < max3; k++ {
					_, err := buffer.Write([]byte(data))
					if err != nil {
						log.Printf("error:%s [%d]", err, id)
					}

				}
				if protecting > 0 {
					mu.Unlock()
				}
			}

		}(i, &buffer)
	}

	for i := 1; i <= max1; i++ {
		<-sign
	}

	data, err := ioutil.ReadAll(&buffer)
	if err != nil {
		log.Fatalf("fatal error:%s", err)
	}
	log.Printf("The contents:%s", data)
}
