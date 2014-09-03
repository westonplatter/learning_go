package main

import "fmt"

func main() {
  // buffer up to 2 inputs
  messages := make(chan string, 2)

  messages <- "buffered"
  messages <- "channel"

  // this next line would break the program
  // messages <- "dropped"

  fmt.Println(<-messages)
  fmt.Println(<-messages)

  // again this next line would break the program
  // fmt.Println(<-messages)
}

