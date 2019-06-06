package main

import "fmt"

type Person struct {
	lastName, firstName string
	iceCream            []string
}

func main() {
	john := Person{
		lastName:  "John",
		firstName: "Snow",
		iceCream:  []string{"Coconut", "Peanut"},
	}
	fmt.Println(john)
}
