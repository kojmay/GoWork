package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
	// for {
	// 	select {
	// 	case j := <-jobs:
	// 		fmt.Println("worker", id, "started  job", j)
	// 		time.Sleep(time.Second)
	// 		fmt.Println("worker", id, "finished job", j)
	// 		results <- j * 2
	// 	case <-time.After(time.Second):
	// 		fmt.Println("...")
	// 	}
	// }

}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, result)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 1; a < numJobs; a++ {
		<-result
	}

}
