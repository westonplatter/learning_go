package main

import "fmt"
import "os"

func main() {
	f := createFile("/tmp/defer.txt")

	// what's the scheduling logic begind defer?
	//
	// http://blog.golang.org/defer-panic-and-recover
	// A defer statement pushes a function call onto a list. The list of saved
	// calls is executed after the surrounding function returns. Defer is commonly
	// used to simplify functions that perform various clean-up actions.
	//
	// https://golang.org/doc/effective_go.html#defer
	// Go's defer statement schedules a function call (the deferred function) to
	// be run immediately before the function executing the defer returns.
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating file")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing file")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing file")
	f.Close()
}
