package main

import "time"
import "fmt"

func main() {
  ticker := time.NewTicker(time.Millisecond * 500)
  go func() {
    for t := range ticker.C {
      fmt.Println("Tick at ", t)
    }
  }()

  time.Sleep(time.Millisecond * 1500)

  // similar to other exercises, manually stop
  // the channel from blocking the main thread
  ticker.Stop()
  fmt.Println("Ticker stopped")
}

