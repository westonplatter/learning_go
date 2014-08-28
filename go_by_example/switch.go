package main 

import "fmt"
import "time"

func main() {
  i := 2

  fmt.Println("write ", i, " as ")
  switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("it's the weekend")
   default:
     fmt.Println("Weekday")
   }

   t := time.Now()
   switch {
   case t.Hour() < 12:
     fmt.Println("am")
   default:
     fmt.Println("pm")
   }
}


