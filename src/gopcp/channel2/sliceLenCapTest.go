package main

import "fmt"

func main() {
	sl := make([]int, 5, 10) //初始化一个长度为5，容量为10的切片，其中的元素值均为切片类型对应的零值
	nsl := sl[1:7]
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(sl), cap(sl), sl)    //长度为5，容量为10，5个元素均为0，未被初始化的元素默认为其零值
	fmt.Printf("len:%d,cap:%d,value:%v\n", len(nsl), cap(nsl), nsl) //长度为6，容量为9，nsl和sl切片的底层数组相同
	//注意：数据的零值是长度为0，容量为0的空数组，而切片的零值是nil，因为切片类型是引用类型，数组类型不是，字典类型也是引用类型，所以零值也是nil
}
