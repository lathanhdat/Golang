package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Printf("%v", x)
	fmt.Printf("\n")
	fmt.Printf("%T", x)
	fmt.Printf("\n")
	x = 42
	fmt.Printf("%v", x)
	fmt.Printf("\n")
}
