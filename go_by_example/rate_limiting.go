package main

import "fmt"
import "time"

func main() {
	// load an int chan with input and close
	requests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		requests <- i
	}
	close(requests)

	// control inflow by blocking on the limiter channel at 200ms intervals
	limiter := time.Tick(time.Millisecond * 200)
	for req := range requests {
		// block thread to limit receiving channel data
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	fmt.Println("-----")

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// inject data into the burstyRequest channel to
	// simulate steady flow system
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
