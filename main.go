package main

import (
	"fmt"
	"go-journey/structs/person"
)

func main () {
	person := &person.Person{ 
		Name: "carl jung", 
		Age: 19, 
	}

	fmt.Println(person.GetName())
	fmt.Println("is", person.GetName(), "under age?: ",person.IsUnderAge())
}