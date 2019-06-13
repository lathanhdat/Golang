package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send5(even, odd, quit)
	receive5(even, odd, quit)
}
func receive5(e, o <-chan int, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even", v)
		case v := <-o:
			fmt.Println("Odd", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("!OK", v, ok)
				return
			} else {
				fmt.Println("OK", v, ok)
				return
			}
		}
	}
}

func send5(e, o chan<- int, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	//Here if q is empty we will receive ok with false value
	q <- 0
	close(q)
}
