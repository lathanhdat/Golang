package main

import "fmt"

func main() {
	y := []string{"Helo", "There"}
	z := []string{"Yello", "Care"}
	x := [][]string{y, z}
	for i, v := range x {
		fmt.Println(i, v)
		for i1, v1 := range v {
			fmt.Println(i1, v1)
		}
	}
	fmt.Println(x)
}
