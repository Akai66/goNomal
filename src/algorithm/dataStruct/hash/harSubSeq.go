package main

import "fmt"

//Input: [1,3,2,2,5,2,3,7]
//Output: 5
//Explanation: The longest harmonious subsequence is [3,2,2,2,3].

//数组中最长和谐子序列的长度
//和谐序列中最大数和最小数之差正好为1，应该注意的是序列的元素不一定是数组的连续元素

func findLHS(arr []int) int {
	max := 0
	imap := map[int]int{}
	for _, v := range arr {
		imap[v] += 1
	}
	for v, n := range imap {
		if n+imap[v+1] > max {
			max = n + imap[v+1]
		}
	}
	return max
}

func main() {
	arr := []int{1, 3, 2, 2, 5, 2, 3, 7}
	fmt.Println(findLHS(arr))
}
