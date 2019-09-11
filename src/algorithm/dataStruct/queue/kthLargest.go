package main

import (
	"algorithm/dataStruct/queue/qdata"
	"fmt"
	"os"
)

//获取数据流中的第k大元素
type KthLargest struct {
	size int //size的含义即为k
	pq   *qdata.MyPriorityQueue
}

func (kl *KthLargest) init(k int, arr []float64) error {
	if k >= 1 && len(arr) >= k-1 {
		kl.size = k
		for _, v := range arr {
			kl.pq.Push(v)
			if kl.pq.Size() > k {
				kl.pq.Pop()
			}
		}
		return nil
	}
	return fmt.Errorf("参数错误")
}

func (kl *KthLargest) add(v float64) float64 {
	kl.pq.Push(v)
	if kl.pq.Size() > kl.size {
		kl.pq.Pop()
	}
	return kl.pq.Peek()
}

func main() {
	kl := &KthLargest{0, &qdata.MyPriorityQueue{Tp: "min"}}
	err := kl.init(3, []float64{1, 6, 5, 4, 3})
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(kl.add(8))
	fmt.Println(kl.add(7))
	fmt.Println(kl.add(9))
	fmt.Println(kl.add(5))
	fmt.Println(kl.add(12))
	fmt.Println(kl.add(10))
}
