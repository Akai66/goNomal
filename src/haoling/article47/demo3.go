package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go startServer1(&wg)

	go startServer2(&wg)

	wg.Wait()
}

func startServer1(wg *sync.WaitGroup) {
	defer wg.Done()
	var httpServer1 http.Server
	httpServer1.Addr = "127.0.0.1:8080"
	// 由于我们没有定制handler，所以这个网络服务对任何请求都只会响应404。
	if err := httpServer1.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("HTTP server 1 closed.")
		} else {
			log.Printf("HTTP server 1 error: %v\n", err)
		}
	}
}

func startServer2(wg *sync.WaitGroup) {
	defer wg.Done()
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/hi" {
			http.NotFound(writer, request)
			return
		}
		name := request.FormValue("name")
		if name == "" {
			fmt.Fprint(writer, "Welcome!")
		} else {
			fmt.Fprintf(writer, "Welcome %s!", name)
		}
	})

	httpServer2 := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mux1,
	}

	if err := httpServer2.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("HTTP server 2 closed.")
		} else {
			log.Printf("HTTP server 2 error: %v\n", err)
		}
	}

}
