package main

import "fmt"

func main() {
	y := returnFunc()
	fmt.Println(y())
}

func returnFunc() func() int {
	x := 10
	return func() int {
		return x
	}
}
