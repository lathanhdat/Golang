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
	denarys := Person{
		lastName:  "Targanya",
		firstName: "Denarys",
		iceCream:  []string{"Cola", "Banana"},
	}

	character := map[string]Person{
		john.lastName:    john,
		denarys.lastName: denarys,
	}

	for i, v := range character {
		fmt.Println(i)
		for i1, v1 := range v.iceCream {
			fmt.Println(i1, v1)
		}
	}

	fmt.Println(character)
}
