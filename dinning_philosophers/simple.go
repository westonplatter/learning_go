package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var workers int = 5
var number_of_simulations = 100

var resources = make([]int, workers)
var load = make([]int, workers)

func random_duration() time.Duration {
	r := rand.Float64() * 20.0
	d := time.Duration(r)
	t := d * time.Millisecond
	return t
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	m := new(sync.Mutex)
	c := make(chan int)

	for i := 0; i < number_of_simulations; i++ {
		go check_and_work(c, *m)
		c <- select_worker(workers)
		time.Sleep(random_duration())
	}

	report()
}

func select_worker(workers int) int {
	// TODO
	// weight selection so distribution is not normalized
	worker := rand.Intn(workers)

	// provide Proportional, Integral, and Derivative correction

	return worker
}

func check_and_work(c chan int, m sync.Mutex) {
	worker := <-c

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
	fmt.Println("- block", number_of_simulations-total_jobs)
	fmt.Println("- completed", total_jobs)
	fmt.Println("- distribution", load_percent)
}
