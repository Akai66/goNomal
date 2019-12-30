package main

import "fmt"

//Given [100, 4, 200, 1, 3, 2],
//The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.

//数组最长连续子序列长度，要求时间复杂度O(n)

func findLCS(arr []int) int {
	max := 0
	imap := map[int]int{}
	for _, v := range arr {
		imap[v] = 1
	}
	for _, v := range arr {
		forward(imap, v)
	}
	for _, n := range imap {
		if n > max {
			max = n
		}
	}
	return max
}

func forward(imap map[int]int, v int) int {
	if _, ok := imap[v]; !ok {
		return 0
	}
	if imap[v] > 1 {
		return imap[v]
	}
	num := forward(imap, v+1) + 1
	imap[v] = num
	return num
}

func main() {
	arr := []int{100, 4, 200, 1, 3, 2, 101}
	fmt.Println(findLCS(arr))
}
