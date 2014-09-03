package main

import "errors"
import "fmt"

// i have no idea what I'm typing
// comments help me with that :)


// method f1
// params, int
// returns, int, error
func f1 (arg int) (int, error) {
  if arg == 42 {
    // construct new error using the go stdlib "error"
    return -1, errors.New("can't work with 42")
  }
  return arg + 3, nil
}


// rather than using the errors.New(), we can
// construct our our struct that impelements the
// execpted interface. here's the struct
type argError struct {
  arg int
  prob string
}

// here's the expected method to implment the interface
func (e *argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// method using our argError struct
func f2(arg int) (int, error) {
  if arg == 42 {

    // Q: why do we have &argError?
    // A: because we can to return the argError value,
    //    not a pointer to the argError
    return -1, &argError{arg, "can't work with 42"}
  }
  return arg + 3, nil
}

func main() {

  // using the the go stdlib errors.New
  for _, i := range[]int{7,42} {
    if r, e := f1(i); e != nil {
      fmt.Println("f1 failed:", e)
    } else {
      fmt.Println("f1 worked:", r)
    }
  }

  // using the argError struct
  for _, i := range []int{7, 42} {
    if r, e := f2(i); e != nil {
      fmt.Println("f2 failed:", e)
    } else {
      fmt.Println("f2 worked:", r)
    }
  }

  // implementing the interface "inline"
  _, e := f2(42)
  if ae, ok := e.(*argError); ok {
    fmt.Println(ae.arg)
    fmt.Println(ae.prob)
  }

}

