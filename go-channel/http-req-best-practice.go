package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			ch <- fmt.Sprintf("%v", i)
		}()

	}
	fmt.Printf("ch: %v\n", ch)
	// for e := range ch {
	// 	fmt.Printf("e: %v\n", e)
	// }

	for i := 0; i < 5; i++ {
		fmt.Println("a", <-ch)
	}

	fmt.Println("-----")

	strTest := ""
	go func() {
		strTest = "1234"
		time.Sleep(3 * time.Second)
	}()
	fmt.Println("test", strTest)

	channel := make(chan int, 1)
	// go func() {
	channel <- 1234
	// time.Sleep(2 * time.Second)
	// }()
	fmt.Println("result", <-channel)

}
