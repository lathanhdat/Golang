package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg4 sync.WaitGroup
var mux sync.Mutex

func main() {
	var count uint64
	count = 0
	times := 100
	wg4.Add(times)
	for i := 0; i < times; i++ {
		go incre4(&count)
	}
	wg4.Wait()
	fmt.Println("count", count)
}

func incre4(number *uint64) {
	atomic.AddUint64(number, 1)
	runtime.Gosched()
	wg4.Done()
}
