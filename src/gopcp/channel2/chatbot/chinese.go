package chatbot

import (
	"fmt"
	"strings"
)

//定义中文机器人结构体类型实现Chatbot接口
type simpleCN struct {
	name string
	talk Talk
}

func NewSimpleCN(name string, talk Talk) Chatbot {
	return &simpleCN{name: name, talk: talk}
}	

func (robot *simpleCN) Name() string {
	return robot.name
}

func (robot *simpleCN) Begin() (string, error) {
	return "请输入你的名字：", nil
}

func (robot *simpleCN) Hello(userName string) string {
	userName = strings.TrimSpace(userName)
	if robot.talk != nil {
		return robot.talk.Hello(userName)
	}
	return fmt.Sprintf("你好，%s！我可以为你做些什么？", userName)
}

func (robot *simpleCN) Talk(heard string) (saying string, end bool, err error) {
	heard = strings.TrimSpace(heard)
	if robot.talk != nil {
		return robot.talk.Talk(heard)
	}
	switch heard {
	case "播放音乐":
		saying = "请欣赏邓紫棋的《倒数》。"
	case "再见", "拜拜":
		saying = "再见！"
		end = true
	default:
		saying = "对不起，我没听懂你说的。"
	}
	return
}

func (robot *simpleCN) ReportError(err error) string {
	return fmt.Sprintf("这里发生了一个错误，%s", err)
}

func (robot *simpleCN) End() (string, error) {
	return "此次聊天结束啦，有空再聊哦。", nil
}
