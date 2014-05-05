package main

import (
	"fmt"
	"math/rand"
)

var workers int = 2
var resources = []int{0}
var load = []int{0, 0}

func main() {
	fmt.Println("start")

	for i := 0; i < 10; i++ {
		worker := choose_worker(workers)
		can_work := resources_available(worker)
		if can_work {
			work(worker)
		}
	}
	fmt.Println("done")

	report()
}

func choose_worker(workers int) int {
	worker := rand.Intn(workers)
	return worker
}

func resources_available(worker int) bool {
	if resources[0] == 0 {
		return true
	} else {
		return false
	}
}

func work(worker int) {
	resources[0] = 1
	load[worker] += 1
	resources[0] = 0
}

func report() {
	total_jobs := 0
	for _, worker_jobs := range load {
		total_jobs += worker_jobs
	}

	load_percent := []float64{0, 0}
	for i := 0; i < len(load); i++ {
		total_jobs := float64(total_jobs)
		worker_jobs := float64(load[i])
		worker_percent := worker_jobs / total_jobs

		load_percent[i] = worker_percent
	}

	fmt.Println("total_jobs", total_jobs)
	fmt.Println("job distribution", load_percent)
}
