package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {

	for i := 0; i < 100; i++ {
		r := rate.Every(4 * time.Second)
		fmt.Println("r", r)
	}
}
