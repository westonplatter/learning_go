package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}

	// built in type specific sorting
	sort.Strings(strs)
	fmt.Println("strings:", strs)

	ints := []int{7, 2, 4}

	// built in type type sorting
	sort.Ints(ints)
	fmt.Println("Ints:    ", ints)

	// more built in helpers
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
}
