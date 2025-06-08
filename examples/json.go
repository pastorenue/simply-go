package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	RelatedTo *Person `json:"related_to"`
	Password  string  `json:"-"`
}

func (p *Person) Serialize() {
	c, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Printf("Result: %s\n", string(c))
}
