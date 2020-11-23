//统计输入中，每个单词出现的频次
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		words[input.Text()]++
	}
	fmt.Printf("word\tcount\n")
	for word, n := range words {
		fmt.Printf("%s\t%d\n", word, n)
	}
}
