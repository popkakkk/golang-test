package main

import (
	"fmt"
	"time"
)

func main() {

	numJobs := 5

	result := make(chan int, numJobs)
	jobs := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go woker(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("a", <-result)
	}

}

func woker(id int, jobs <-chan int, result chan<- int) {
	fmt.Println("init worker", id)
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(2 * time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
	fmt.Println("end worker", id)
}
