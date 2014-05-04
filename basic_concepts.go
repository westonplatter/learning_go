package main

import (
	"fmt"
	"math"
)

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

//var i, j int = 1,2
//var c, python, java = true, false, "no!"

func add(x int, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	// v is not acessible here, since it was declared
	// within the if statement's scope

	return lim
}

func Sqrt(x float64) float64 {
	// z = z - ((z^2 - x)/2z)

	z := float64(1)

	for i := 1.0; i < 10; i++ {
		z = z - ((math.Pow(z, 2) - x) / 2 * z)
		fmt.Println(i, z)
	}

	return z
}

type Vertex struct {
	X, Y int
}

func main() {

	// UT8 by default
	fmt.Println("Hello, 世界")

	// use imported packages
	fmt.Println("Now you have the %g problems.", math.Nextafter(2, 3))

	// use defined functions
	fmt.Println("2 + 2 =", add(2, 2))

	// functions with multiple arguments of same type of declare type once
	fmt.Println("5 - 3 =", subtract(5, 3))

	// functions can return 2 values
	a, b := swap("up", "down")
	fmt.Println("swapping 'up down'", a, b)

	// functions implicitly return named variables
	fmt.Println(split(17))

	// variable assignment
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)

	// constants
	fmt.Println(Pi)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))

	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// continue loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// while loop
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// if statement
	if 2 > 1 {
		fmt.Println("2 is greater than 1")
	}

	// if short statement
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	// build sqrt to show loops and functions
	fmt.Println(Sqrt(2))

	// instantiate a struct
	fmt.Println(Vertex{1,2})

	// struct values are mutable
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)

	// struct pointers
	// pointer arithmetic not available
	// http://www.cplusplus.com/doc/tutorial/pointers/#arithmetics
	p := Vertex{1,2}
	q := &p
	q.X = 1e9
	fmt.Println(p)

	// new function
	// "allocates zero type value and returns pointer to it
	h := new(Vertex)
	fmt.Println(h)
	h.X, h.Y = 11, 9
	fmt.Println(h)

	// arrays, which cannot be resized
	var aa [2]string
	aa[0] = "Hello"
	aa[1] = "World"
	fmt.Println(aa[0], aa[1])
	fmt.Println(aa)

	// slices - these are important
	ka := []int{2,3,5,7,11,13}
	fmt.Println("ka == ", ka)
	for i := 0; i < len(ka); i++ {
		fmt.Printf("ka[%d] == %d\n", i, ka[i])
	}

	// slicing slices
	// the new slice points to the same array
	kk := ka[0:3]
	kk[0] = 100
	fmt.Println("kk has mutable access to ka, ka[0] = ",ka[0])
	// slicing actions		
	fmt.Println("ka[1:4] == ", ka[1:4])

}
