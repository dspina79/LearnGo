package main

import "fmt"

const PI float32 = 3.14159

func add(x int, y int) int {
	return x + y
}

func greet(name string) {
	fmt.Println("Hello, ", name)
}

func areaOfCircle(radius float32) float32 {
	return PI * (radius * radius)
}

func main() {
	var x, y int = 5, 6
	var z = add(x, y)
	var nom = "Dave"
	var radius float32 = 2.0
	greet(nom)
	fmt.Println("The sum of ", x, " and ", y, " is ", z)
	fmt.Println("A cirlce with a radius of ", radius, " is ", areaOfCircle(radius))
}
