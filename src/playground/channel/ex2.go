package main

import "fmt"

func main() {
	rc := make(chan int)
	go send(rc)
	receive(rc) //We will block here by not using go

}

func send(c chan<- int) {
	c <- 42
}
func receive(c <-chan int) {
	fmt.Println(<-c)
}
