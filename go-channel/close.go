package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan string, 5)
	done := make(chan bool, 1)
	go func() {

		time.Sleep(time.Second * 3)
		for {
			r, more := <-jobs
			if more {
				fmt.Println("received job", r)
			} else {
				fmt.Println("received all job")
				done <- true
				return
			}
		}

	}()

	for i := 0; i < 3; i++ {

		time.Sleep(time.Second * 1)
		jobs <- fmt.Sprintf("%v", i)
	}

	close(jobs)

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
	<-done

}
