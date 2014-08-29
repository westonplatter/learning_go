package main

import "fmt"

func plus(a, b int) int {
  return a + b
}

func main() {
  result := plus(1, 2)
  fmt.Println("1+2=", result)
}

