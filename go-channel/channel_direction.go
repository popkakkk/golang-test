package main

import (
	"fmt"
)

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	msg := "Pop patipan"

	ping(pings, msg)
	pong(pings, pongs)

	fmt.Println("messge : ", <-pongs)

}

func ping(pigns chan<- string, msg string) { // receive only
	pigns <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // send only , receive only
	pongs <- <-pings
}
