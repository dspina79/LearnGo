package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) float32 {
	fx := float32(x)
	fy := float32(y)

	return fx / fy
}

func main() {
	var x int = 10
	var y int = 15

	fmt.Println(x, " + ", y, "=", add(x, y))
	fmt.Println(x, " x ", y, "=", multiply(x, y))
	fmt.Println(x, " / ", y, "=", divide(x, y))
}
