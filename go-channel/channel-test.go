package main

import (
	"fmt"
	"time"
)

func main() {

	cha1 := make(chan int, 2)
	cha2 := make(chan int, 2)

	cha1 <- 1
	cha1 <- 2

	cha2 <- 1

	go func() {
		time.Sleep(time.Hour)
	}()

	fmt.Println("cha1", <-cha1)
	fmt.Println("cha1", <-cha1) // dead lock for wait cha1
	fmt.Println("cha2", <-cha2)
	fmt.Println("cha2", <-cha2)
}
