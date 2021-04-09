package main

import "fmt"

func dbl() func(x int) int {
	return func(x int) int {
		return x * 2
	}
}

func main() {
	x := 4
	d := dbl()
	fmt.Println("x is starting as", x)

	fmt.Println("Double x is", d(x))
	fmt.Println("Double-double x is", d(d(x)))
	fmt.Println("Double-double-double x is", d(d(d(x))))
}
