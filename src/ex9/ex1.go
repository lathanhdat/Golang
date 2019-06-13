package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)
	go roundtine1()
	go roundtine2()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Done")

}

func roundtine1() {
	fmt.Println("Hello from roundtine1")
	wg.Done()
}

func roundtine2() {
	fmt.Println("Hello from roundtine2")
	wg.Done()
}
