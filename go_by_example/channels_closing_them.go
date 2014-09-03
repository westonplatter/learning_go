package main

import "fmt"

func main() {
  jobs := make(chan int, 5)
  done := make(chan bool)

  go func() {
    for {
      // i is the int declared when chan was initialized
      // more is bool telling us if channel has been closed
      i, more := <-jobs

      if more {
        fmt.Println("received job", i)
      } else {
        fmt.Println("received all jobs")
        done <- true

        // why did the example have a return here?
        // return
      }
    }
  }()

  for j := 1; j <= 3; j ++ {
    jobs <- j
    fmt.Println("send job", j)
  }
  close(jobs)
  fmt.Println("sent all jobs")

  // channel blocks program from ending unti it receives
  // the input from the anonymous function
  <-done
}


