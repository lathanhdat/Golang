package main

import "fmt"

func main() {
	defer fooDefer()
	fooNormal()
}
func fooDefer() {
	fmt.Println("Defer")
}
func fooNormal() {
	fmt.Println("Normal")
}
