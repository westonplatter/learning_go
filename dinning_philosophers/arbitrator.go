package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// determining factors for simulation

var workers int = 10
var number_of_simulations = 100
var random_duration_range float64 = 5.0

// variables calculated from simulation variables

var resources = make([]int, workers)
var load = make([]int, workers)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	m := new(sync.Mutex)
	c := make(chan int)

	for i := 0; i < number_of_simulations; i++ {
		time.Sleep(1.0 * time.Millisecond)
		go check_and_work(c, *m)
		c <- select_worker(workers)
	}

	report()
}

func check_and_work(c chan int, m sync.Mutex) {
	worker := <-c

	// lock and unlock the Mutex to ensure 1 and only 1 "goroutinue"
	// accesses the shared `resources` variable at a time.

	m.Lock()
	can_work := resources_available(worker)

	if can_work {
		reserve(worker)
		m.Unlock()

		work(worker)

		m.Lock()
		free(worker)
	}

	m.Unlock()
}

func select_worker(workers int) int {

	// intentionally weight the distribution of random numbers so first
	// half of worker units receive all the work completion requests
	max := workers - int(workers/2)
	worker := rand.Intn(max)

	if completed_more_than(0.35) {
		// after 35% of all potential jobs have been completed,
		// re-select the worker to be a worker with less
		// than or equal to the min number of jobs completed
		worker = min_worker()
	}

	return worker
}

func min_worker() int {
	min := load[0]
	worker := 0

	for i := 0; i < len(load); i++ {
		if load[i] < min {
			min = load[i]
			worker = i
		}
	}

	return worker
}

func completed_more_than(percent float64) bool {
	return percent_done() > percent
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

func reserve(worker int) {
	left := left(worker)
	right := right(worker)

	resources[left] = 1
	resources[right] = 1
}

func free(worker int) {
	left := left(worker)
	right := right(worker)

	resources[left] = 0
	resources[right] = 0
}

func work(worker int) {
	time.Sleep(random_duration())
	load[worker] += 1
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

func random_duration() time.Duration {
	r := rand.Float64() * random_duration_range
	d := time.Duration(r)
	t := d * time.Millisecond
	return t
}

func report() {
	total_jobs := 0
	for _, worker_jobs := range load {
		total_jobs += worker_jobs
	}

	load_percent := make([]float64, workers)

	for i := 0; i < len(load); i++ {
		total_jobs := float64(total_jobs)
		worker_jobs := float64(load[i])
		worker_percent := worker_jobs / total_jobs

		load_percent[i] = worker_percent
	}

	fmt.Println("Jobs", number_of_simulations)
	fmt.Println("- standard_deviation", standard_deviation())
	fmt.Println("- blocked", number_of_simulations-total_jobs)
	fmt.Println("- completed", total_jobs)
	fmt.Println("- distribution", load_percent)
}

// simple data/metrics functions

func percent_done() float64 {
	t := total_completed()
	num := float64(number_of_simulations)
	return t / num
}

func total_completed() float64 {
	total := 0.0

	for _, value := range load {
		total += float64(value)
	}

	return total
}

func standard_deviation() float64 {
	total_completed := total_completed()
	numerator := 0.0

	for _, value := range load {
		worker_jobs := float64(value)
		worker_percent := worker_jobs / total_completed
		numerator += worker_percent
	}

	return math.Sqrt(numerator / total_completed)
}
