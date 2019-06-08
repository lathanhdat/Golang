package main

import "fmt"

func main() {
	x := closesure()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func closesure() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}
