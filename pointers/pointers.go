package main

import "fmt"
// * -> pointer
// & -> memory address

// this function copy the arguments, and does not affect the value
func increment(num int) {
	num++
}

func Increment(num *int) {
	// Never modifies the argument directly
	*num++
}

func main() {
	value := 10
	fmt.Println("Value before the increment", value, &value)
	Increment(&value)
	fmt.Println("Value after the increment", value)

	// new()
	pointer := new(int) // pointer int initialized in zero
	*pointer = 20
	fmt.Println("initial value with new: ", *pointer)
}