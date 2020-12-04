package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func main() {
	//reader接口
	res, err := readFrom(os.Stdin, 8)
	if err != nil {
		fmt.Printf("readFrom error:%v\n", err)
		return
	}
	fmt.Printf("%s\n", res)

	//ReaderAt接口
	r := strings.NewReader("Go语言是一门很棒的语言")
	res, _ = readAt(r, 2)
	fmt.Printf("%s,%d\n", res, len(res))

	//WriterAt接口
	//新建一个文件
	tmpDir := os.TempDir()
	filePath := filepath.Join(tmpDir, "qafile01.txt")
	fmt.Println(filePath)
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	f.WriteString("Golong中国社区--这里是多余")
	data := make([]byte, 10)
	n, err := f.Read(data) //即使之前没有读过，但是此时该文件的读取计数已经到了文件末尾，所以此时读取响应的err=EOF
	fmt.Printf("data:%s,n:%d,err:%v\n", data, n, err)
	f.Close()

	//ReadFrom接口
	//重新打开文件读取
	f, _ = os.Open(filePath)
	w := bufio.NewWriter(os.Stdout)
	num, _ := w.ReadFrom(f)
	w.Flush()
	fmt.Printf("\nread %d\n", num)
	f.Close()

	//WriteTo接口
	wr := bytes.NewReader([]byte("Go语言中文网\n"))
	resn, err := wr.WriteTo(os.Stdout)
	fmt.Printf("write %d\n", resn)

	//Seeker接口
	r = strings.NewReader("Go语言中文网")
	r.Seek(-6, io.SeekEnd)
	c, size, _ := r.ReadRune()
	fmt.Printf("char:%q,size:%d\n", c, size)

	//Closer接口
	f, err = os.Open(filePath)
	//一般在错误校验之后，使用defer，因为如果文件打开失败后，f是nil，调用Close方法会出错
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	defer f.Close()

	//ByteReader和ByteWriter接口
	var ch byte
	fmt.Scanf("%c\n", &ch)
	buffer := new(bytes.Buffer)
	err = buffer.WriteByte(ch)
	if err != nil {
		fmt.Printf("write error:%v\n", err)
	}
	fmt.Println("写入成功，准备读取一个字节")
	newCh, _ := buffer.ReadByte()
	fmt.Printf("读取的字节:%c\n", newCh)
	fmt.Println("读取一个字节成功")

	//ByteScanner接口
	buffer = bytes.NewBuffer([]byte("ab"))
	err = buffer.UnreadByte() //不能直接使用UnreadByte,在使用之前必须先调用ReadByte
	fmt.Printf("error:%v\n", err)
	ch, _ = buffer.ReadByte()
	fmt.Printf("%q\n", ch)
	err = buffer.UnreadByte()
	fmt.Printf("first unread :%v\n", err)
	err = buffer.UnreadByte()
	fmt.Printf("second unread :%v\n", err)
	ch, _ = buffer.ReadByte()
	fmt.Printf("%q\n", ch)

	//关于strings.Index方法
	fmt.Println(strings.Index("Go语言中文网", "中文"))
	fmt.Println(Utf8Index("Go语言中文网", "中文"))

	//SectionReader类型
	r = strings.NewReader("abcdefg")
	sr := io.NewSectionReader(r, 1, 4)
	tmp := make([]byte, 2)
	for {
		n, err := sr.Read(tmp)
		fmt.Printf("data:%s,n:%d,err:%v\n", tmp, n, err)
		if err != nil {
			break
		}
	}

	//LimitedReader，限制最多读取多少个字节
	fmt.Println()
	r = strings.NewReader("testLimitedReader")
	lr := io.LimitReader(r, 10)
	for {
		_, err := lr.Read(tmp)
		if err != nil {
			break
		}
		fmt.Printf("%s", tmp)
	}

}

func readFrom(r io.Reader, num int) ([]byte, error) {
	data := make([]byte, num)
	n, err := r.Read(data)
	if n > 0 {
		return data[:n], nil
	}
	return data, err
}

func readAt(r io.ReaderAt, offset int64) ([]byte, error) {
	data := make([]byte, 6)
	n, err := r.ReadAt(data, offset)
	if n > 0 {
		return data[:n], nil
	}
	return data, err
}

func writeAt(f *os.File, offset int64, s string) (n int, err error) {
	n, err = f.WriteAt([]byte(s), offset)
	return
}

func Utf8Index(str, substr string) int {
	i := strings.Index(str, substr)
	if i < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:i])
}
