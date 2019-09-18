package main

import (
	"algorithm/dataStruct/queue/qdata"
	"fmt"
	"os"
)

/**
 * 滑动窗口的最大值
 * 给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。
 */

//方法一，利用双端队列实现
//用双端队列来存储数组元素的索引
//1.如果新来的值比队列尾部的数小，那就追加到后面，因为它可能在前面的最大值划出窗口后成为最大值
//2.如果新来的值比尾部的大，那就删掉尾部，不断和新尾部进行比较，直到队列为空或者小于新的尾部，最后追加到队列后面
//3.如果追加的值的索引跟队列头部的值的索引超过窗口大小，那就删掉头部的值
//4.每次队列的头都是滑动窗口中值最大的
func maxInWindows1(arr []float64, size int) ([]float64, error) {
	if size > 0 && size <= len(arr) {
		var resArr []float64
		mdq := new(qdata.MyDeque)
		for i := 0; i < len(arr); i++ {
			if !mdq.Empty() {
				if i >= mdq.Peek().(int)+size {
					mdq.Pop()
				}
				for !mdq.Empty() && arr[i] > arr[mdq.GetLast().(int)] {
					mdq.RemoveLast()
				}
			}
			mdq.Push(i)
			if i+1 >= size {
				resArr = append(resArr, arr[mdq.Peek().(int)])
			}
		}
		return resArr, nil
	}
	return nil, fmt.Errorf("参数错误")
}

//方法二，利用优先级队列实现
//最大优先级队列
//1.队列元素数到达窗口大小，就获取堆顶元素
//2.队列元素超过窗口大小，就删除不在窗口中元素
func maxInWindows2(arr []float64, size int) ([]float64, error) {
	if size > 0 && size <= len(arr) {
		mpq := &qdata.MyPriorityQueue{Tp: "max"}
		var resArr []float64
		for i := 0; i < len(arr); i++ {
			mpq.Push(arr[i])
			if i+1 == size {
				resArr = append(resArr, mpq.Peek())
			}
			if i+1 > size {
				mpq.Remove(arr[i-size])
				resArr = append(resArr, mpq.Peek())
			}
		}
		return resArr, nil
	}
	return nil, fmt.Errorf("参数错误")
}

func main() {
	arr := []float64{4, 6, 7, 9, 10, 11, 12, 7, 14, 13}
	res, err := maxInWindows1(arr, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(res)
	res, err = maxInWindows2(arr, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(res)
}
