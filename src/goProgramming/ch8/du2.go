//遍历目录，统计目录下文件总大小
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

var sema = make(chan struct{},20)

func main() {
	flag.Parse()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(10 * time.Millisecond) //每间隔多少时间会向该chan发送一个消息
	}
	roots := flag.Args()
	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSizes, &wg)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
    //限制打开文件个数
    sema<- struct{}{}   //chan满了之后，会被阻塞，防止同时打开的文件个数过多
    defer func(){<-sema}()        //处理完后，接收消息，腾出一个缓存空间
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du:%v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files,%.2f GB\n", nfiles, float64(nbytes)/1e9)
}
