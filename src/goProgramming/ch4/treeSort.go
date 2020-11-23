//先构造搜索二叉树，然后使用中序遍历，实现排序
package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	arr := []int{4, 3, 5, 2, 4, 6, 8, 7}
	sort(arr)
	fmt.Println(arr)
}

func sort(arr []int) {
	var root *tree
	//构造搜索二叉树
	for _, value := range arr {
		root = add(root, value)
	}
	//中序遍历
	appendValue(arr[:0], root) //arr和原始arr共用底层数组，所以这里的修改会直接反映在原始参数变量上
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if v < t.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}

func appendValue(a []int, t *tree) []int {
	if t != nil {
		a = appendValue(a, t.left)
		a = append(a, t.value)
		a = appendValue(a, t.right)
	}
	return a
}
