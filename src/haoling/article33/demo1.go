package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

//代表存放数据块缓冲区的临时对象池
var bufPool sync.Pool

var delimiter = byte('\n')

type Buffer interface {
	Delimiter() byte
	Write(contents string) (err error)
	Read() (contents string, err error)
	Free()
}

type myBuffer struct {
	buf       bytes.Buffer
	delimiter byte
}

func (b *myBuffer) Delimiter() byte {
	return b.delimiter
}

func (b *myBuffer) Write(contents string) (err error) {
	if _, err = b.buf.WriteString(contents); err != nil {
		return
	}
	return b.buf.WriteByte(b.delimiter)
}

func (b *myBuffer) Read() (contents string, err error) {
	return b.buf.ReadString(b.delimiter)
}

func (b *myBuffer) Free() {
	//用完之后再put回临时对象池，保证临时对象池中总有比较充足的临时对象
	bufPool.Put(b)
}

func init() {
	bufPool = sync.Pool{
		//指定New字段，保证没有获取到临时对象时，最终调用New字段对应的方法创建一个
		New: func() interface{} {
			return &myBuffer{delimiter: delimiter}
		},
	}
}

func GetBuff() Buffer {
	return bufPool.Get().(Buffer)
}

func main() {
	buf := GetBuff()
	defer buf.Free()
	buf.Write("A Pool is a set of temporary objects that may be individually saved and retrieved.")
	buf.Write("A Pool is safe for use by multiple goroutines simultaneously.")
	buf.Write("A Pool must not be copied after first use.")

	fmt.Println("The data blocks in buffer:")

	for {
		block, err := buf.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("unexpected error:%s", err))
		}
		fmt.Print(block)
	}
}
