package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type personJSON struct {
	FirstParsed string `json:"First"`
	LastParsed  string `json:"Last"`
	AgeParsed   string `json:Age`
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   23,
	}

	p2 := person{
		First: "James",
		Last:  "Bamd",
		Age:   29,
	}

	Agent := []person{p1, p2}
	fmt.Println(Agent)

	bs, err := json.Marshal(Agent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))

	AgentParsed := []personJSON{}
	error1 := json.Unmarshal(bs, &AgentParsed)
	if error1 != nil {
		fmt.Println(err)
	}
	fmt.Println(AgentParsed)
}

//Don't know why but we need to user upper case for first character of field's name
