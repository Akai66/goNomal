//修改charcount函数，使用unicode.IsLetter等相关函数，统计字母和数字不同字符的数量
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	allChars := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int //统计utf8编码字节数分类
	letters := make(map[rune]int)
	numbers := make(map[rune]int)
	invalid := 0 //记录无效的UTF8字符

	input := bufio.NewReader(os.Stdin)
	for {
		r, n, err := input.ReadRune() //返回UTF8解码后的rune,以及该字符使用UTF8编码后占用的字节数，最后是err
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		allChars[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			letters[r]++
		}
		if unicode.IsNumber(r) {
			numbers[r]++
		}
	}
	fmt.Printf("char\tcount\n")
	for c, n := range allChars {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nletter\tcount\n")
	for l, n := range letters {
		fmt.Printf("%q\t%d\n", l, n)
	}

	fmt.Printf("\nnumber\tcount\n")
	for i, n := range numbers {
		fmt.Printf("%q\t%d\n", i, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}

	fmt.Printf("\n%d invalid UTF8 chars \n", invalid)
}
