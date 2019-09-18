package qdata

import (
	"fmt"
)

//优先级队列，每次出队列的是最大值或最小值，利用二叉堆实现
type MyPriorityQueue struct {
	size int //记录队列长度,不能以切片长度为准，因为在push的时候是给size+1指定位置赋值，pop的时候只更新size的值，并不删除切片中的元素
	arr  []float64
	Tp   string //"max"表示最大优先级队列，"min"表示最小优先级队列
}

//入队列
func (mpq *MyPriorityQueue) Push(v float64) {
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
func (mpq *MyPriorityQueue) Pop() float64 {
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
	mpq.downAdjust(0)
	return maxV
}

//删除指定元素
func (mpq *MyPriorityQueue) Remove(v float64) {
	if mpq.size <= 0 {
		fmt.Println("Remove error：队列为空")
		return
	}
	removeIndex := -1
	for i := 0; i < mpq.size; i++ {
		if mpq.arr[i] == v {
			removeIndex = i
			break
		}
	}
	if removeIndex < 0 {
		fmt.Println("Remove error：被删除的元素不存在")
		return
	}
	//将尾部元素赋值给被删除的位置
	mpq.arr[removeIndex] = mpq.arr[mpq.size-1]
	mpq.size -= 1
	//二叉堆进行"下沉"操作
	mpq.downAdjust(removeIndex)
}

//获取队列首部
func (mpq *MyPriorityQueue) Peek() float64 {
	if mpq.size <= 0 {
		fmt.Println("Peek error：队列为空")
		return 0
	}
	return mpq.arr[0]
}

func (mpq *MyPriorityQueue) Size() int {
	return mpq.size
}

//二叉堆"上浮"操作
func (mpq *MyPriorityQueue) upAdjust() {
	childIndex := mpq.size - 1
	parentIndex := (childIndex - 1) / 2
	temp := mpq.arr[childIndex]
	if mpq.Tp == "max" {
		for childIndex > 0 && temp > mpq.arr[parentIndex] {
			//优化点：无需真正交换，单方向赋值即可
			mpq.arr[childIndex] = mpq.arr[parentIndex]
			childIndex = parentIndex
			parentIndex = (parentIndex - 1) / 2
		}
	} else {
		for childIndex > 0 && temp < mpq.arr[parentIndex] {
			//优化点：无需真正交换，单方向赋值即可
			mpq.arr[childIndex] = mpq.arr[parentIndex]
			childIndex = parentIndex
			parentIndex = (parentIndex - 1) / 2
		}
	}

	mpq.arr[childIndex] = temp
}

//二叉堆"下沉"操作
func (mpq *MyPriorityQueue) downAdjust(index int) {
	parentIndex := index
	childIndex := index*2 + 1
	temp := mpq.arr[parentIndex]
	if mpq.Tp == "max" {
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
	} else {
		for childIndex < mpq.size {
			if childIndex+1 < mpq.size && mpq.arr[childIndex+1] < mpq.arr[childIndex] {
				childIndex += 1
			}
			if temp < mpq.arr[childIndex] {
				break
			}
			//优化点：无需真正交换，单方向赋值即可
			mpq.arr[parentIndex] = mpq.arr[childIndex]
			parentIndex = childIndex
			childIndex = childIndex*2 + 1
		}
	}

	mpq.arr[parentIndex] = temp
}

//对切片进行扩容
func (mpq *MyPriorityQueue) resize() {
	//对切片进行扩容
	mpq.arr = append(mpq.arr, make([]float64, 20)...)
}
