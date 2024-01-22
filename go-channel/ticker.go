package main

import (
	"fmt"
	"time"
)

func main() {

	tiker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool, 3)

	go func() {
		time.Sleep(2 * time.Second)
		tiker.Stop()
		done <- true
	}()

	selectChan(tiker.C, done)

	go func() {
		time.Sleep(3 * time.Second)
		done <- false
	}()

	ok := <-done
	fmt.Println("ok", ok)

}

func selectChan(t <-chan time.Time, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		case resp := <-t:
			fmt.Println("ticker 2 :", resp)
		}
	}
}
