package main

import (
	"fmt"
	"github.com/zoumo/goset" //github上实现的set
)

//题目描述：求两个数组的交集，结果去重

//思路：运用set，set会自动去重

func intersection(arr1, arr2 []int) (isIst bool, istSet goset.Set) {
	set1 := goset.NewSet()
	istSet = goset.NewSet()
	for _, v := range arr1 {
		set1.Add(v)
	}
	for _, v := range arr2 {
		if set1.Contains(v) {
			istSet.Add(v)
		}
	}
	if istSet.Len() > 0 {
		isIst = true
	}
	return isIst, istSet
}

//题目描述：求两个数组的交集，结果不去重

//运用hash表

func intersection2(arr1, arr2 []int) (isIst bool, istArr []int) {
	imap := map[int]int{}
	istArr = []int{}
	for _, v := range arr1 {
		imap[v] += 1
	}
	for _, v := range arr2 {
		if imap[v] > 0 {
			istArr = append(istArr, v)
			imap[v] -= 1
		}
	}
	if len(istArr) > 0 {
		isIst = true
	}
	return isIst, istArr
}

func main() {
	arr1 := []int{1, 2, 3, 3, 2}
	arr2 := []int{1, 1, 3, 3, 3, 4, 4}
	fmt.Println(intersection(arr1, arr2))
	fmt.Println(intersection2(arr1, arr2))
}
