package main

import "fmt"
import "time"

func worker(done chan bool) {
  fmt.Print("working ...")
  time.Sleep(time.Second)
  fmt.Println("done")

  // send confirmation that worker action is done
  done <- true
}

func main() {
  done := make(chan bool, 1)
  go worker(done)

  // go blocks the program from ending until we hear
  // back from the the channel (sent to the worker)
  <-done
}

