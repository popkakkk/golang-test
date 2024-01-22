package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	a, b := <-c2

	go func() {
		time.Sleep(time.Second * 4)
		c1 <- "pop patipan"
	}()

	select {
	case resp := <-c1:
		fmt.Println("resp", resp)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "mieng thanyarat"
	}()

	select {
	case resp := <-c2:
		fmt.Println("resp", resp)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout 2")
	}
}
