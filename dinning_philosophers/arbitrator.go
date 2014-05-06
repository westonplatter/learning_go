/*

Copyright (c) 2014, Weston Platter, BSD-3

Problem
Dinning Philosophers
http://en.wikipedia.org/wiki/Dining_philosophers_problem

Solution
Self Correcting Arbitrator

The solution implements an "Arbitrator" resource management
pattern. When a worker wants to work, it checks if resources
are available, if they are the worker reserves the resources,
completes its work, and then releases the resources. If
resources are not available, the worker unit stops work
execution.

The resource availability state is tracked in a global
variable, `resources`. Workers lock the mutex before reading
and writing to the `resources` variable. This is the arbitrator
pattern.

The solution is also self correcting. By default, the
simulation favors workers that are in the first half of the
worker population (see the `SelectWorker` function). In order
to ensure work is evenly distributed across the worker
population, the `SelectWorker` function proportionally
reassigns work to lesser worked worker units. While the
solution provides a simple proportional system control,
solution, a full PID control system could be implemented
to account for unpredictable and rapidly changing work loads.

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// determining factors for simulation

var workers int = 5
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
		go ChecKAndWork(c, *m)
		c <- SelectWorker(workers)
	}

	Report()
}

func ChecKAndWork(c chan int, m sync.Mutex) {
	worker := <-c

	// lock and unlock the Mutex to ensure 1 and only 1
	// "goroutinue" accesses the shared `resources`
	// variable at a time.

	m.Lock()
	can_work := ResourcesAvailable(worker)

	if can_work {
		Reserve(worker)
		m.Unlock()

		Work(worker)

		m.Lock()
		Free(worker)
	}

	m.Unlock()
}

func SelectWorker(workers int) int {

	// intentionally weight the distribution of random numbers so first
	// half of worker units receive all the work completion requests
	max := workers - int(workers/2)
	worker := rand.Intn(max)

	if CompletedMoreThan(0.35) {
		// after 35% of all potential jobs have been completed,
		// re-select the worker to be a worker with less
		// than or equal to the min number of jobs completed

		// the choice of 35% is abitrary. This is offset from
		// 0% percent so that there's a sufficient data to use
		// to make work reassignment decisions
		worker = MinWorker(load)
	}

	return worker
}

func MinWorker(load []int) int {
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

func CompletedMoreThan(percent float64) bool {
	return PercentDone() > percent
}

func ResourcesAvailable(worker int) bool {
	left := Left(worker)
	right := Right(worker, workers)

	if resources[left] == 0 && resources[right] == 0 {
		return true
	} else {
		return false
	}
}

func Reserve(worker int) {
	left := Left(worker)
	right := Right(worker, workers)

	resources[left] = 1
	resources[right] = 1
}

func Free(worker int) {
	left := Left(worker)
	right := Right(worker, workers)

	resources[left] = 0
	resources[right] = 0
}

func Work(worker int) {
	time.Sleep(RandomDuration())
	load[worker] += 1
}

func Left(i int) int {
	return i
}

func Right(i int, workers int) int {
	right := i - 1
	if right == -1 {
		right = workers - 1
	}
	return right
}

func RandomDuration() time.Duration {
	r := rand.Float64() * random_duration_range
	d := time.Duration(r)
	t := d * time.Millisecond
	return t
}

func Report() {
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
	fmt.Println("- blocked", number_of_simulations-total_jobs)
	fmt.Println("- completed", total_jobs)
	
	fmt.Println("\nworker percents")
	
	for i := 0; i < len(load_percent); i++ {
		fmt.Printf("- %v = %.2f\n", i, load_percent[i])
	}
}

// simple data/metrics functions

func PercentDone() float64 {
	t := TotalCompleted()
	num := float64(number_of_simulations)
	return t / num
}

func TotalCompleted() float64 {
	total := 0.0

	for _, value := range load {
		total += float64(value)
	}

	return total
}

