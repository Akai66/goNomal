package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mailbox uint8

	var lock sync.RWMutex

	//代表专用于发信的条件变量
	sendCond := sync.NewCond(&lock)

	//代表专用于收信的条件变量
	recvCond := sync.NewCond(lock.RLocker())
	//用于传递演示完成的信号，防止主goroutine
	sign := make(chan struct{}, 2)
	max := 5

	//用于发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			log.Printf("sender [%d]:the mailbox is empty.", i)
			mailbox = 1
			log.Printf("sender [%d]:the letter has been sent.", i)
			lock.Unlock()
			recvCond.Signal()
		}

	}(max)

	//用于收信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			log.Printf("receiver [%d]:the mailbox is full.", i)
			mailbox = 0
			log.Printf("receiver [%d]:the letter has been received.", i)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<-sign
	<-sign
}
