package main

import (
	"log"
	"sync"
	"time"
)

type counter struct {
	num uint
	mu  sync.RWMutex
}

func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

func (c *counter) add(incr uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += incr
	return c.num
}

func main() {
	c := counter{}
	count(&c)
}

func count(c *counter) {
	sign := make(chan struct{}, 3)

	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 20; i++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter:%d [%d-%d]", c.number(), 1, i)
		}
	}()

	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 20; i++ {
			time.Sleep(time.Millisecond * 300)
			log.Printf("The number in counter:%d [%d-%d]", c.number(), 2, i)
		}
	}()

	<-sign
	<-sign
	<-sign
}
