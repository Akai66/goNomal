package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	//如何通过原子操作，减少int32或uint32的值
	num := int32(18)
	fmt.Println(num)
	atomic.AddInt32(&num, -3)
	fmt.Println(atomic.LoadInt32(&num))

	num2 := uint32(18)
	fmt.Println(num2)
	delta := int32(-3)
	atomic.AddUint32(&num2, uint32(delta))
	fmt.Println(atomic.LoadUint32(&num2))

	//简易的自旋锁
	forAndCAS1()

	//简易的,更加宽松的互斥锁模拟
	forAndCAS2()

}

func forAndCAS1() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number :%d\n", num)
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for {
			time.Sleep(time.Millisecond * 500)
			newNum := atomic.AddInt32(&num, 2)
			fmt.Printf("The number :%d\n", newNum)
			if newNum == 10 {
				break
			}
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for {
			if atomic.CompareAndSwapInt32(&num, 10, 0) {
				fmt.Println("The number has gone to zero")
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	<-sign
	<-sign
}

func forAndCAS2() {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	max := int32(20)

	go func(id int, max int32) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; ; i++ {
			curNum := atomic.LoadInt32(&num)
			if curNum >= max {
				break
			}
			newNum := curNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, curNum, newNum) {
				fmt.Printf("The number:%d [%d-%d]\n", newNum, id, i)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(1, max)

	go func(id int, max int32) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 0; ; i++ {
			curNum := atomic.LoadInt32(&num)
			if curNum >= max {
				break
			}
			newNum := curNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, curNum, newNum) {
				fmt.Printf("The number:%d [%d-%d]\n", newNum, id, i)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(2, max)

	<-sign
	<-sign
}
