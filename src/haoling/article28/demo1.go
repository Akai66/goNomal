package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	//代表信箱，0表示信箱是空的，1表示信箱是满的
	var mailbox uint8

	var lock sync.Mutex

	sendCond := sync.NewCond(&lock)

	recvCond := sync.NewCond(&lock)

	send := func(id, index int) {
		lock.Lock()
		for mailbox == 1 {
			sendCond.Wait()
		}
		log.Printf("sender [%d-%d]:the mailbox is empty", id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]:the letter has been sent", id, index)
		lock.Unlock()
		recvCond.Broadcast() //确定有多个接收的goroutine
	}

	recv := func(id, index int) {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		log.Printf("receiver [%d-%d]:the mailbox is full", id, index)
		mailbox = 0
		log.Printf("receiver [%d-%d]:the letter has been received", id, index)
		lock.Unlock()
		sendCond.Signal() //确定只有一个发送的goroutine
	}

	sign := make(chan struct{}, 3)

	max := 6

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 500)
			send(id, i)
		}
	}(1, max)

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, i)
		}
	}(1, max/2)

	go func(id, max int) {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, i)
		}
	}(2, max/2)

	<-sign
	<-sign
	<-sign
}
