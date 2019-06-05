package main

import "fmt"

func main() {
	x := [5]int{234, 2534, 646, 2235, 6456}
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T \n", x)
}
