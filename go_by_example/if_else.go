package main

import "fmt"

func main() {
  if 7%2 == 0 {
    fmt.Println("7 is even")
  } else {
    fmt.Println("7 is correctly odd")
  }

  even := 8%4 == 0
  if even {
    fmt.Println("eight is divisible by 4")
  }

  if num := 9; num < 0 {
    fmt.Println("9 is greater than 0")
  } else if num < 10 {
    fmt.Println("9 is less than 10")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}


