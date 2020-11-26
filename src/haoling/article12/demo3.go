package main

import "fmt"

//go中函数参数的传递，都是"浅拷贝"，只是看参数本身是值类型还是引用类型，值类型函数内部对参数的修改不会改变原参，引用类型则会影响到原参
//函数的结果在赋值时，也会进行拷贝，都是"浅拷贝"，同参数

func main() {
	arr := [3]int{1, 2, 3}
	arr2 := modify(arr)
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", arr2)

	//复杂数组，外层是数组，内层嵌套切片，外层值类型，内层引用类型
	complexArr := [3][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	//只改变外层，不影响原参
	complexArr1 := modifyOut(complexArr)
	fmt.Printf("%v\n", complexArr)
	fmt.Printf("%v\n", complexArr1)

	//改变内层，影响原参
	complexArr2 := modifyInner(complexArr)
	fmt.Printf("%v\n", complexArr)
	fmt.Printf("%v\n", complexArr2)

}

func modify(arr [3]int) [3]int {
	arr[0] = 2
	return arr
}

func modifyOut(arr [3][]int) [3][]int {
	arr[0] = []int{2, 2, 2}
	return arr
}

func modifyInner(arr [3][]int) [3][]int {
	arr[0][0] = 2
	return arr
}
