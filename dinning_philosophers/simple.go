package main

import (
	"fmt"
	"math/rand"
	"time"
)

var workers int = 5
var resources = []int{0, 0, 0, 0, 0}
var load = []int{0, 0, 0, 0, 0}

var time_interval_request time.Duration = 2.0 * time.Millisecond
var time_interval_work time.Duration = 3.0 * time.Millisecond

func main() {
	for i := 0; i < 10; i++ {
		worker := choose_worker(workers)
		can_work := resources_available(worker)
		if can_work {
			go async_work(worker)
		} else {
			fmt.Println(worker, "- blocked")
		}

		time.Sleep(time_interval_request)
	}
	fmt.Println("done")

	report()
}

func choose_worker(workers int) int {
	worker := rand.Intn(workers)
	return worker
}

func resources_available(worker int) bool {
	left := left(worker)
	right := right(worker)

	if resources[left] == 0 && resources[right] == 0 {
		return true
	} else {
		return false
	}
}

func async_work(worker int) {
	fmt.Println(worker, "+ working")

	left := left(worker)
	right := right(worker)

	// reserve resources
	resources[left] = 1
	resources[right] = 1
	// work
	time.Sleep(time_interval_work)
	// record metrics
	load[worker] += 1
	// put resources back
	resources[left] = 0
	resources[right] = 0
}

func left(i int) int {
	return i
}

func right(i int) int {
	right := i - 1
	if right == -1 {
		right = workers - 1
	}
	return right
}

func report() {
	total_jobs := 0
	for _, worker_jobs := range load {
		total_jobs += worker_jobs
	}

	load_percent := []float64{0, 0, 0, 0, 0}
	for i := 0; i < len(load); i++ {
		total_jobs := float64(total_jobs)
		worker_jobs := float64(load[i])
		worker_percent := worker_jobs / total_jobs

		load_percent[i] = worker_percent
	}

	fmt.Println("total_jobs", total_jobs)
	fmt.Println("job distribution", load_percent)
}
