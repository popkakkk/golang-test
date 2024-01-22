package main

import (
	"fmt"
	"time"
)

func main() {

	cha1 := make(chan string, 1)
	cha2 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		cha1 <- "PP"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		cha2 <- "MM"
	}()

	for i := 0; i < 2; i++ {
		select {
		case resp := <-cha1:
			fmt.Println("cha1 :", resp)
		case resp := <-cha2:
			fmt.Println("cha2:", resp)
		}
	}
}
