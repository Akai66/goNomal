package main

import "fmt"

//题目描述
//所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。
//
//编写一个函数来查找 DNA 分子中所有出现超过一次的 10 个字母长的序列（子串）。

func findRepeatedDnaSequences(str string) ([]string, map[int32]int32) {
	if len(str) < 10 {
		fmt.Println("参数异常")
	}
	var mask, cur int32 = 0x7ffffff, 0
	seqHash := map[int32]int32{}
	resSeqs := []string{}
	for i := 0; i < 9; i++ {
		cur = (cur << 3) | (int32(str[i]) & 7)   //数字7二进制为三个1，与7进行&操作是为了将该数字前面的位数全部置为0，仅保留后3位数
		//每次左移三位，是为了将补位的3个0和&操作后的三位数进行|操作，将cur的最后三位置为str[i]的后三位
	}
	for i := 9; i < len(str); i++ {
		cur = ((cur & mask) << 3) | (int32(str[i]) & 7)  //与mask进行&操作是为了保留cur的后27位，并将其它位置为0
		if _, ok := seqHash[cur]; ok {
			if seqHash[cur] == 1 {
				resSeqs = append(resSeqs, str[i-9:i+1])
			}
			seqHash[cur] += 1
		} else {
			seqHash[cur] = 1
		}
	}
	return resSeqs, seqHash
}

func main() {
	str := "ACGTTAGTACGTACGTTAGTACGTACGACGTTTTAGTACGTACGATAGTACGTACGTTAGTACGTACGACGTTTTAGTACTAGTACGTACGACGTTTTA"
	fmt.Println(findRepeatedDnaSequences(str))
}
