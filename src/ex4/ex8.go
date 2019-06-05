package main

import "fmt"

func main() {
	x := map[string][]string{
		"RecordA": []string{"bond-james", "Martis"},
		"RecordB": []string{"money_pey", "Letera"},
		"RecordC": []string{"Icecream", "limen"},
	}
	for i, v := range x {
		fmt.Println(i)
		for i1, v1 := range v {
			fmt.Println(i1, v1)
		}
	}
	fmt.Println(x)

}
