package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var (
	count int64
)

func increment() error {
	count = count + 1
	return nil
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	increment()
	fmt.Fprint(w, count)
	log.Println("Number of Requests", atomic.LoadInt64(&count))
	atomic.AddInt64(&count, 0)
}
