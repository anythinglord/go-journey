package main

import "fmt"

func main () {
	var array = [5]int{1,2,3,4,5}
	var array2 = [...]int{1,2,3,4,5}
	array[1] = 15 
	fmt.Println(array, array2)
	for i:=0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for index, value := range array {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	var matrix = [3][3] int{{1,0,0}, {0,1,0}, {0,0,1}}
	fmt.Println(matrix)
}