package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now().UTC()

	fmt.Println("n", n)

	loc, _ := time.LoadLocation("Asia/Bangkok")

	new := n.In(loc)
	fmt.Println("new", new)
}
