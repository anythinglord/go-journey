package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero :(")
	}
	result := a / b
	return result, nil
}

func printNames(names...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// clousure example
func counter() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

func main() {
	result, error := division(10, 3)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(result)
	count := counter()
	fmt.Println("count is:", count())
}