package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("now :", time.Now())
	time.Sleep(3 * time.Second)
	t1 := <-time.After(3 * time.Second)

	fmt.Println("t1 :", t1)
}
