package main

import (
	"fmt"
	"math/rand"
	"time"
)

var workers int = 2
var resources = []int{0}

func main() {
	fmt.Println("starting")
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		worker := choose_worker(workers)
		can_work := resources_available(worker)
		if can_work {
			work(worker)
		}
	}
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
	fmt.Println("worker", worker, "starting work")
	resources[0] = 1
	time.Sleep(2 * time.Second)
	resources[0] = 0
}
