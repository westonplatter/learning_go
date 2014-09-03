package main

import "fmt"

func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  // synchronously call method
  f("direct")

  // call method via gorountine
  // executes concurrently
  go f("gorountine")


  go func(msg string) {
    fmt.Println(msg)
  }("going")

  var input string
  fmt.Scanln(&input)
  fmt.Println("done")
}

