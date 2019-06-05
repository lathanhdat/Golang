package main

import "fmt"

func main() {
	i := 1995
	for {
		if i >= 1995 {
			fmt.Println(i)
		}
		if i >= 2019 {
			break
		}
		i++
	}
}
