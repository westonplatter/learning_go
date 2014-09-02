package main

import "fmt"

type person struct {
  name string
  age  int
}

func main() {
  fmt.Println(person{"Bob", 20})

  fmt.Println(&person{"pointer to struct", 10})

  s := person{name: "jill", age: 32}
  fmt.Println(s.name)

  spointer := &s
  fmt.Println("automatically dereferenced - ", spointer.name)

  s.age = 51
  fmt.Println(s.age)
}

