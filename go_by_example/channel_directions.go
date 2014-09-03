package main

import "fmt"

// pings, sending channel
func ping(pings chan<- string, msg string) {

  // msg argument is sent into the pings argument
  pings <- msg
}

// pings, receiving channel
// pongs, sending channel
func pong(pings <-chan string, pongs chan<- string) {

  // gets input from the pings argument
  msg := <-pings

  // sends the message into the pongs arugment
  pongs <- msg
}

func main() {
  pings := make(chan string, 1)
  pongs := make(chan string, 1)

  // send before receiving end is setup.
  // we can do this since channel is buffered 1x
  ping(pings, "passed message")

  // setup the middle part
  pong(pings, pongs)

  // pritn out whatever is in the system
  fmt.Println(<-pongs)

  // notice there has to be parity in the buffering of all
  // channel connections
}

