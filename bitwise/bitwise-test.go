package main

import "fmt"

type Day int

const (
	Mon Day = 1 << iota
	Tues
	Wed
	Thurs
)

func main() {
	fmt.Println("Mon", Mon)
	fmt.Println("Tues", Tues)
	fmt.Println("Wed", Wed)
	fmt.Println("Thurs", Thurs)

	fmt.Print("Mon|Tues", Mon|Tues)
}
