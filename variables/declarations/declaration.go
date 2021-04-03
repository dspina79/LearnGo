package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func greet(name string) {
	fmt.Println("Hello, ", name)
}

func main() {
	var x, y int = 5, 6
	var z = add(x, y)
	var nom = "Dave"
	greet(nom)
	fmt.Println("The sum of ", x, " and ", y, " is ", z)
}
