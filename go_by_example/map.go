package main 

import "fmt"

func main() {

  // make MAP[key-type]value-type
  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1:", v1)

  fmt.Println("len:", len(m))

  delete(m, "k2")
  fmt.Println("map:", m)

  // handle error if key is not present in map
  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  n := map[string]int{"foo": 1, "bar":2}
  fmt.Println("map:", n)
}

