package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(500 * time.Millisecond)
	go clientGo(1)
	wg.Wait()
}

func serverGo() {
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen Error:%s", err)
		return
	}
	defer listener.Close()
	printServerLog("Got listener for the server.(local address:%s)", listener.Addr())
	//获取监听器后不断等待客户端新的连接请求
	for {
		conn, err := listener.Accept() //阻塞直到有新的客户端连接到来
		if err != nil {
			printServerLog("Accept error:%s", err)
		}
		printServerLog("Established a connection with a client application.(remote address:%s)", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func clientGo(id int) {
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printClientLog(id, "Dial Error: %s", err)
		return
	}
	defer func() {
		conn.Close()
		wg.Done()
	}()
	printClientLog(id, "Connected to server. (remote address: %s, local address: %s)", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	requestTime := 5
	for i := 0; i < requestTime; i++ {
		req := rand.Int63()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientLog(id, "Write error:%s", err)
			continue
		}
		printClientLog(id, "Send request (written %d bytes):%d", n, req)
	}

	for j := 0; j < requestTime; j++ {
		strRes, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientLog(id, "The connection is closed by another side.")
			} else {
				printClientLog(id, "Read error:%s", err)
			}
			break
		}
		printClientLog(id, "Received response: %s", strRes)
	}
}

//服务端处理客户端请求
func handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
		wg.Done()
	}()
	for {
		//读取客户端发送的消息
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerLog("The connection is closed by another side.")
			} else {
				printServerLog("Read error:%s", err)
			}
			break
		}
		printServerLog("Received request:%s", strReq)
		//响应客户端消息
		//判断字符串是否可以转为64位整型
		intReq, err := strToInt64(strReq)
		if err != nil {
			n, _ := write(conn, err.Error())
			printServerLog("Send error message (written %d bytes):%s", n, err)
			continue
		}
		//求立方根
		floatResp := cbrt(intReq)
		respMsg := fmt.Sprintf("The cube root of %d is %f", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerLog("Write error:%s", err)
		}
		printServerLog("Send response (write %d bytes):%s", n, respMsg)
	}
}

//读取消息
func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

//写消息
func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

//字符串转int64
func strToInt64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("\"%s\" is not integer", str)
	}
	if num > math.MaxInt64 || num < math.MinInt64 {
		return 0, fmt.Errorf("%d is not 64-bit integer", num)
	}
	return int64(num), err
}

//求立方根
func cbrt(param int64) float64 {
	return math.Cbrt(float64(param))
}

//打印日志
func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]:%s", role, sn, fmt.Sprintf(format, args...))
}

//打印服务端日志
func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args...)
}

//打印客户端日志
func printClientLog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args...)
}
