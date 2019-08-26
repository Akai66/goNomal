package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopcp/channel2/chatbot"
	"os"
	"runtime/debug"
)

var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.cn", "The chatbot's name for dialogue.")
}

func main() {
	flag.Parse()
	chatbot.Register(chatbot.NewSimpleCN("simple.cn", nil))
	myChatbot := chatbot.Get(chatbotName)
	if myChatbot == nil {
		err := fmt.Errorf("Fatal error: Unsupported chatbot named %s\n", chatbotName)
		checkError(nil, err, true)
	}
	inputReader := bufio.NewReader(os.Stdin)
	begin, err := myChatbot.Begin()
	checkError(myChatbot, err, true)
	fmt.Println(begin)
	input, err := inputReader.ReadString('\n')
	checkError(nil, err, true)
	fmt.Println(myChatbot.Hello(input[:len(input)-1]))
	for {
		input, err := inputReader.ReadString('\n')
		if checkError(nil, err, false) {
			continue
		}
		saying, end, err := myChatbot.Talk(input)
		if checkError(myChatbot, err, false) {
			continue
		}
		fmt.Println(saying)
		if end {
			endStr, err := myChatbot.End()
			if !checkError(myChatbot, err, false) {
				fmt.Println(endStr)
			}
			os.Exit(0)
		}

	}
}

func checkError(robot chatbot.Chatbot, err error, exit bool) bool {
	if err == nil {
		return false
	} else {
		if robot != nil {
			fmt.Println(robot.ReportError(err))
		} else {
			fmt.Println(err)
		}
		if exit == true {
			debug.PrintStack()
			os.Exit(1)
		}
		return true
	}
}
