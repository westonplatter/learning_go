package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// basic printing
	fmt.Printf("%v\n", p)

	// print struct key/value names
	fmt.Printf("%+v\n", p)

	// print syntax representation of value and key/value names
	fmt.Printf("%#v\n", p)

	// print value type
	fmt.Printf("%T\n", p)

	// print boolean values in plain manner
	fmt.Printf("%t\n", true)

	// print digits in base-10 manner
	fmt.Printf("%d\n", 123)

	// print digits in binary or base-2 manner
	fmt.Printf("%b\n", 14)

	// print char of given value
	fmt.Printf("%c\n", 33)

	// print hex encoding
	fmt.Printf("%x\n", 456)

	// print float values
	fmt.Printf("%f\n", 78.5)

	// print scientific notation
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// print basic string
	fmt.Printf("%s\n", "normal string")

	// print string with 2x quotes
	fmt.Printf("%q\n", "\"more strings\"")

	// print string as hex
	fmt.Printf("%x\n", "string")

	// print pointer
	fmt.Printf("%p\n", &p)

	// print controlling width and precision
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// print floating points the same way
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// print strings controlling width and distance
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// print strings with alignment
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// print strings without seeing output
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// use io.Writers to print output to StandardOut
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
