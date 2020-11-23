//典型的并发循环处理
package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	sizes := make(chan int)
	var wg sync.WaitGroup
	files := []string{"image1", "image2", "image3"}
	for _,f := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			_, err := imageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			sizes <- 2
		}(f)
	}

	//这一步很关键，需要新起一个goroutine来wait，close
	go func() {
		wg.Wait()
		close(sizes)
	}()

	total := 0
	for size := range sizes {
		total += size
	}

	fmt.Println(total)
}

//伪代码，主要关注典型的并发处理写法
func imageFile(f string) (string, error) {
	return f, nil
}
