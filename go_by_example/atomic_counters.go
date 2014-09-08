package main

import "time"
import "fmt"
import "sync/atomic"
import "runtime"

func main() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {

				// why are we using the memory address, "&ops" ?
				//
				// http://golang.org/pkg/sync/atomic/
				// the atomic package provies low-level memory primatives useful for
				// implementing sync algorithms
				//
				// this is a special chase, and communication across gorountines should
				// almost always be handled through channels not shared memory access.
				//
				atomic.AddUint64(&ops, 1)

				// I don't know why we need this line
				//
				// runtime.Gosched()
				// Gosched yields the processor, allowing other goroutines to run.
				//
				// If I don't add this line, the process doesn't finish. So if it's not
				// included, the multiple gorountines are not able to execute hence the
				// process jam.
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)

	fmt.Println("printing ops also works:", ops)
}
