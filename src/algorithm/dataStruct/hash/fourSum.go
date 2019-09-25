package main

import "fmt"

//题目描述
//给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，使得 A[i] + B[j] + C[k] + D[l] = 0。
//
//为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。所有整数的范围在 -2^28 到 2^28- 1 之间，最终结果不会超过 2^31 - 1 。

//解题思路
//只是求符合条件的元组个数
//与 Two Sum 极其类似，使用哈希表来解决问题。
//
//把 A 和 B 的两两之和都求出来，在哈希表中建立两数之和与其出现次数之间的映射；
//
//遍历 C 和 D 中任意两个数之和，只要看哈希表存不存在这两数之和的相反数就行了。

func fourSum(A, B, C, D []int) int {
	if len(A) <= 0 || len(B) <= 0 || len(C) <= 0 || len(D) <= 0 {
		fmt.Print("参数错误")
		return 0
	}
	res := 0
	mp := map[int]int{}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			mp[A[i]+B[j]] += 1
		}
	}
	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			if _, ok := mp[-(C[i] + D[j])]; ok {
				res += mp[-(C[i] + D[j])]
			}
		}
	}
	return res
}

func main() {
	A := []int{1, 2, 1}
	B := []int{-1, -1, 2}
	C := []int{-1, 1, 3}
	D := []int{0, 2, -2}
	fmt.Print(fourSum(A, B, C, D))
}
