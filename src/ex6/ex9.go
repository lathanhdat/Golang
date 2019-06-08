package main

import "fmt"

func main() {
	fmt.Println(callBack(number))
}

func callBack(x func(numb1 int) int) int {
	numb := 1
	y := x(numb)
	return y
}

func number(numb1 int) int {
	return numb1 + 10
}
