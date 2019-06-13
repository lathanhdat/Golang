package main

import "fmt"

func main() {
	rc := make(chan int, 100)
	go send(rc)
	for y := range rc {
		fmt.Println(y)
	}
	fmt.Println("Exit")
}

func send(c chan<- int) {
	for _i := 0; _i < 100; _i++ {
		c <- _i
	}
	// close(c)
}
