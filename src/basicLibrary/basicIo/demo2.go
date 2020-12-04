package main

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//PipeReader和PipeWriter类型
	pReader, pWriter := io.Pipe()
	wg.Add(2)
	go PipeWrite(pWriter)
	go PipeRead(pReader)
	wg.Wait()
}

func PipeWrite(w *io.PipeWriter) {
	defer wg.Done()
	data := []byte("Go语言中文网")
	for i := 0; i < 3; i++ {
		n, err := w.Write(data)
		if err != nil {
			fmt.Printf("write error:%v\n", err)
			break
		}
		fmt.Printf("Writer:写入%d个字节\n", n)
	}
	w.CloseWithError(errors.New("写入端关闭"))
}

func PipeRead(r *io.PipeReader) {
	defer wg.Done()
	data := make([]byte, 128)
	//读之前先sleep一会,让写入端阻塞
	fmt.Println("Reader:读取端开始sleep2秒")
	time.Sleep(time.Second * 2)
	fmt.Println("Reader:读取端开始读取数据")
	for {
		n, err := r.Read(data)
		if err != nil {
			fmt.Printf("Reader:error:%v\n", err)
			break
		}
		fmt.Printf("Reader:读取数据:%s,共%d个字节\n", data[:n], n)
	}
}
