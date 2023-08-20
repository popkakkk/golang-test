package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	x := now.AddDate(0, 0, 3)

	fmt.Println("x", x)

}
