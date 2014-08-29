package main 

import "fmt"

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}

func main() {
  i := 1
  fmt.Println("intiail:", i)

  // no change because we're operating on the set variable,
  // not the data structure persisted in memory
  // see next
  zeroval(i)
  fmt.Println("zeroval:", i)

  // we're passing in the "i" memory address by passing in "&i"
  zeroptr(&i)
  fmt.Println("zeroptr:", i)
}
