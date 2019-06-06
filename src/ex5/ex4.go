package main

import "fmt"

func main() {
	john := struct {
		lastName  string
		firstName string
		friend    []string
		food      map[string]int
	}{
		lastName:  "Snow",
		firstName: "John",
		friend:    []string{"Denarys", "Tyrion"},
		food: map[string]int{
			"Burger": 55,
			"Hotdog": 79,
		},
	}
	fmt.Println(john)
	for k, v := range john.food {
		fmt.Println(k, v)
	}
	for i, v := range john.friend {
		fmt.Println(i, v)
	}
}
