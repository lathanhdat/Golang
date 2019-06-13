package main

import "fmt"

type Person struct {
	name string
}

func (s *Person) speak() {
	fmt.Println(s.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	john := Person{
		name: "John Snow",
	}

	saySomething(&john)

}
