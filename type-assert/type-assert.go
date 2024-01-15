package main

import (
	"fmt"
)

func main() {
	// var value interface{} = []int{1, 2, 3}
	var value2 interface{} = nil

	slice, ok := value2.([]int)
	if !ok {
		fmt.Println("The value is not a slice of integers.")
		return
	}

	fmt.Println(slice) // prints: [1 2 3]
}
