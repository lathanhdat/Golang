package main

import "fmt"

func main() {
	x := make([]string, 55, 60)
	fmt.Println(cap(x))
	fmt.Println(len(x))
}
