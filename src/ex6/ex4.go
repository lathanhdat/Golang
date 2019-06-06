package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (s Person) speak() {
	fmt.Println(s.first, s.last, s.age)
}
func main() {
	john := Person{
		first: "John",
		last:  "Snow",
		age:   10,
	}
	john.speak()
}
