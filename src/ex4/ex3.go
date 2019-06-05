package main

import "fmt"

func main() {
	x := []int{1, 23, 535, 2, 234, 53, 634, 234, 345, 23}
	fmt.Println(x[:2])
	fmt.Println(x[4:])
}
