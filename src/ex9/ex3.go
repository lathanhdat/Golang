package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg3 sync.WaitGroup

// var count int

func main() {
	count := 0
	times := 100
	wg3.Add(times)
	for i := 0; i < times; i++ {
		go incre(&count)
	}
	wg3.Wait()
	fmt.Println("count", count)
}

func incre(number *int) {
	numbStore := *number
	runtime.Gosched()
	numbStore++
	*number = numbStore
	wg3.Done()
}
