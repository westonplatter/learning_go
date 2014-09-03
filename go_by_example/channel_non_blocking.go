package main

import "fmt"

func main() {
  messages := make(chan string)
  signals := make(chan bool)

  // non-blocking receive
  // ie, if nothing in the messages channel, default case
  // question - shouldn't messages have a buffer?
  select {
  case msg := <-messages:
    fmt.Println("received message", msg)
  default:
    fmt.Println("no message received")
  }

  // non-blocking send
  // immediately send the msg var
  msg := "hi"
  select {
  case messages <-msg:
    fmt.Println("sent message", msg)
  default:
    fmt.Println("no message sent")
  }

  // non-blocking receive for multiple channels
  select {
  case msg := <-messages:
    fmt.Println("receieved message", msg)
  case sig := <-signals:
    fmt.Println("receieved signal", sig)
  default:
    fmt.Println("no activity")
  }
}


