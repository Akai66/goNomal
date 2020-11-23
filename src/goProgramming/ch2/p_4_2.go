//根据命令行参数指定是否用sha512,默认sha256,对标准输入生成摘要
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var s512 = flag.Bool("s512", false, "is use sha512")

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "bye" {
			os.Exit(0)
		}
		if *s512 {
			fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
		} else {
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		}
	}
}
