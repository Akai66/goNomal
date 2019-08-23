//接口类型的定义，实现，继承
package main

import (
	"bufio"
	"fmt"
	"os"
)

//定义接口类型
//go语言的接口类型用于定义一组行为，其中每个行为都由一个方法声明表示，接口类型的方法声明只有方法签名而没有方法体，方法签名包括且仅包括方法的名称，参数列表和结果列表
type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

//接口类型的实现
//只要一个数据类型的方法集合中包含该接口声明的所有方法，那么它一定是该接口的实现类型，一个接口类型的变量可以被赋予任何该接口实现类型的值
type myTalk string

func (talk *myTalk) Hello(userName string) string {
	ret := fmt.Sprintf("你好! %s", userName)
	return ret
}

func (talk *myTalk) Talk(heard string) (saying string, end bool, err error) {
	if heard == "音乐" {
		saying = "播放邓紫棋的《倒数》"
	} else if heard == "退出" {
		saying = "再见!"
		end = true
	} else {
		saying = "不好意思，我听不懂你的意思"
	}
	return
}

//go的数据结构之间并不存在继承关系，接口类型之间也是如此，不过一个接口类型的声明中可以嵌入任意其它接口类型，通俗的说，一组行为可以包含其它的行为组，且数量不限
type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() (string, error)
}

//myTalk实现Chatbot接口
func (talk *myTalk) Name() string {
	return "机器人名称"
}

func (talk *myTalk) Begin() (str string, err error) {
	str = "开始聊天"
	return
}

func (talk *myTalk) ReportError(err error) (str string) {
	return
}

func (talk *myTalk) End() (str string, err error) {
	str = "结束聊天"
	return
}

//与myTalk关联的全部是指针方法，所以myTalk类型并不是Talk接口的实现类型，*myTalk类型才是
func main() {
	var tk Talk = new(myTalk) //内建函数new的功能是创建一个指定类型的值，并返回指向该值的指针
	_, ok := tk.(*myTalk)
	fmt.Println(ok)
	var cb Chatbot = new(myTalk)
	_, ok = cb.(*myTalk)
	fmt.Println(ok)
	begin, _ := cb.Begin()
	fmt.Println(begin)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字：")
	input, _ := inputReader.ReadString('\n')
	fmt.Println(tk.Hello(input[:len(input)-1]))
	fmt.Println("请尝试与我对话：")
	for {
		input, _ = inputReader.ReadString('\n')
		saying, end, _ := tk.Talk(input[:len(input)-1])
		fmt.Println(saying)
		if end == true {
			end, _ := cb.End()
			fmt.Println(end)
			break
		}
	}
}
