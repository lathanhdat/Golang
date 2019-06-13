package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		channel <- 42
	}()
	v, ok := <-channel
	if !ok {
		fmt.Println("!OK", v, ok)
	} else {
		fmt.Println("OK", v, ok)
	}
}
