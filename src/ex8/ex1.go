package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

type user []person

func main() {
	person1 := person{
		Name: "Tdat1",
		Age:  14,
	}
	person2 := person{
		Name: "Tdat2",
		Age:  24,
	}

	user1 := user{person1, person2}
	bs, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
