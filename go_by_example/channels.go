package main

import "fmt"

func main() {
  messages := make(chan string)

  // start gorountine with anonymous function
  // sending "ping" into messages channel
  go func() { messages <- "ping"}()

  msg := <-messages
  fmt.Println(msg)
}
