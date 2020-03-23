package main

import (
	"fmt"
)

//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
//说明：
//
//拆分时可以重复使用字典中的单词。
//
//你可以假设字典中没有重复的单词。
//在这里，只需要去定义一个数组 bool[] isSplitArr，其中第 i 位 isSplitArr[i] 表示待拆分字符串从第 0 位到第 i-1 位是否可以被成功地拆分。
//
//然后分别计算每一位是否可以被成功地拆分。

func wordBreak(str string, wordDict map[string]int) bool {
	//参数校验
	strLen := len(str)
	wordDictLen := len(wordDict)
	if strLen <= 0 || wordDictLen <= 0 {
		fmt.Print("params error")
		return false
	}
	max := 0
	for word, _ := range wordDict {
		wordLen := len(word)
		if wordLen > max {
			max = wordLen
		}
	}

	isSplitArr := make([]bool, strLen+1)
	isSplitArr[0] = true
	for i := 1; i <= strLen; i++ {
		for j := i - 1; j >= 0 && i-j <= max; j-- {
			subStr := str[j:i]
			if _, ok := wordDict[subStr]; ok && isSplitArr[j] {
				isSplitArr[i] = true
				break
			}
		}
	}
	return isSplitArr[strLen]
}

func main() {
	str := "abcdefa"
	wordDict := map[string]int{"b": 1, "cd": 1, "e": 1, "f": 1, "a": 1}
	fmt.Println(wordBreak(str, wordDict))
}
