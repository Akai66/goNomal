package main

import "fmt"

//最大优先级队列，每次出队列的是最大值，利用最大二叉堆实现
type myPriorityQueue struct {
	size int //记录队列长度,不能以切片长度为准，因为在push的时候是给size+1指定位置赋值，pop的时候只更新size的值，并不删除切片中的元素
	arr  []float64
}

//入队列
func (mpq *myPriorityQueue) Push(v float64) {
	//入队列，二叉堆进行"上浮"操作
	//扩容操作
	if mpq.size >= len(mpq.arr) {
		mpq.resize()
	}
	mpq.arr[mpq.size] = v
	mpq.size += 1
	mpq.upAdjust()

}

//出对列
func (mpq *myPriorityQueue) Pop() float64 {
	if mpq.size <= 0 {
		fmt.Println("Pop error：队列为空")
		return 0
	}
	//将数组首元素取出
	maxV := mpq.arr[0]
	//将尾部元素赋值给首部
	mpq.arr[0] = mpq.arr[mpq.size-1]
	mpq.size -= 1
	//二叉堆进行"下沉"操作
	mpq.downAdjust()
	return maxV
}

//获取队列首部
func (mpq *myPriorityQueue) Peek() float64 {
	if mpq.size <= 0 {
		fmt.Println("Peek error：队列为空")
		return 0
	}
	return mpq.arr[0]
}

//二叉堆"上浮"操作
func (mpq *myPriorityQueue) upAdjust() {
	childIndex := mpq.size - 1
	parentIndex := (childIndex - 1) / 2
	temp := mpq.arr[childIndex]
	for childIndex > 0 && temp > mpq.arr[parentIndex] {
		//优化点：无需真正交换，单方向赋值即可
		mpq.arr[childIndex] = mpq.arr[parentIndex]
		childIndex = parentIndex
		parentIndex = (parentIndex - 1) / 2
	}
	mpq.arr[childIndex] = temp
}

//二叉堆"下沉"操作
func (mpq *myPriorityQueue) downAdjust() {
	parentIndex := 0
	childIndex := 1
	temp := mpq.arr[parentIndex]
	for childIndex < mpq.size {
		if childIndex+1 < mpq.size && mpq.arr[childIndex+1] > mpq.arr[childIndex] {
			childIndex += 1
		}
		if temp > mpq.arr[childIndex] {
			break
		}
		//优化点：无需真正交换，单方向赋值即可
		mpq.arr[parentIndex] = mpq.arr[childIndex]
		parentIndex = childIndex
		childIndex = childIndex*2 + 1
	}
	mpq.arr[parentIndex] = temp
}

//对切片进行扩容
func (mpq *myPriorityQueue) resize() {
	//对切片进行扩容
	mpq.arr = append(mpq.arr, make([]float64, len(mpq.arr))...)
}

func main() {
	mpq := myPriorityQueue{0, make([]float64, 5)}
	mpq.Push(1)
	fmt.Println(mpq.Peek())
	mpq.Push(3)
	fmt.Println(mpq.Pop())
	mpq.Push(7)
	fmt.Println(mpq.Peek())
	mpq.Push(5)
	fmt.Println(mpq.Pop())
	fmt.Println(mpq.Pop())
	mpq.Push(4)
	fmt.Println(mpq.Peek())
	mpq.Push(7)
	fmt.Println(mpq.Pop())
}
