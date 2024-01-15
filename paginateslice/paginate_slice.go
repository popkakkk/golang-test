package main

import (
	"fmt"
)

func paginateSlice(data []int, offset int, limit int) []int {
	if offset > len(data) {
		return nil
	}

	end := offset + limit
	if end > len(data) {
		end = len(data)
	}

	return data[offset:end]
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	offset := 2
	limit := 5
	result := paginateSlice(data, offset, limit)
	fmt.Println(result) // [3 4 5 6 7]
}
