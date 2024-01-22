package main

import (
	"fmt"
	"time"
)

func main() {

	name := make(chan string, 10)
	go worker(name)

	fmt.Println("result of name :", <-name)
}

func worker(name chan string) {
	go process(name)
}

func process(v chan string) {
	fmt.Println("procress running")
	time.Sleep(time.Second * 3)
	v <- "is your gf name"
}
