package main

import "fmt"

func main () {
	dict := map[string] int {
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	fmt.Println(dict["key1"])
}