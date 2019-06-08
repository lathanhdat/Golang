package main

import "fmt"

type person struct {
	name string
}

func changeMe(s *person, name string) {
	// (*s).name = name
	s.name = name
}

func main() {
	john := person{
		name: "John Snow",
	}
	changeMe(&john, "John Know Nothing")
	fmt.Println(john)
}
