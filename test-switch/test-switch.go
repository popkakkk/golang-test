package main

import "fmt"

func main() {

	myFunc(1)
}

func myFunc(a int) {
	fmt.Println("start")
	switch a {
	case 1:
	case 2:

	}

	fmt.Println("end")
}
