package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			work(i)
		}()
	}
	wg.Wait()
}
func work(id int) {
	fmt.Println("Worker starting", id)
	time.Sleep(time.Second)
	fmt.Println("Worker done", id)
}
