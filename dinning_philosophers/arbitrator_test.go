package main

import (
	"fmt"
	"testing"
)

// func TestCheckAndWork(t *testing.T) {
// }

// func TestSelectWorker(t *testing.T) {
// }

func TestMinWorker(t *testing.T) {
	var load = []int{3, 0, 5}

	min := MinWorker(load)

	if min != 1 {
		t.Fail()
	}
}

// func TestCompletedMoreThan(t *testing.T) {
// }

// func TestResourcesAvailable(t *testing.T) {
// }

// func TestReserve(t *testing.T) {
// }

// func TestFree(t *testing.T) {
// }

// func TestWork(t *testing.T) {
// }

func TestLeft(t *testing.T) {
	if l := Left(1); 1 != l {
		t.Fail()
	}
}

func TestRight(t *testing.T) {
	if r := Right(1, 5); 0 != r {
		t.Fail()
	}

	if r := Right(0, 5); 4 != r {
		fmt.Println(r)
		t.Fail()
	}
}

// func TestRandomDuration(t *testing.T) {
// }
