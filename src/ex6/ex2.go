package main

import "fmt"

func main() {
	numbs := []int{12, 53, 24, 6, 2, 3}
	sum := sum(numbs...)
	fmt.Println(sum)
	sum = sum1(numbs)
	fmt.Println(sum)
}

func sum(numbs ...int) int {
	sum := 0
	for _, v := range numbs {
		sum += v
	}
	return sum
}

func sum1(numbs []int) int {
	sum := 0
	for _, v := range numbs {
		sum += v
	}
	return sum
}
