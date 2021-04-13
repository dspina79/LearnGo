package main

import (
	"errors"
	"fmt"
)

func divde(x, y int) (float32, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return float32(x) / float32(y), nil
}

func main() {
	var x int
	var y int

	fmt.Print("Enter a numerator: ")
	fmt.Scanf("%d", &x)
	fmt.Print("Enter a denominator: ")
	fmt.Scanf("%d", &y)

	result, err := divde(x, y)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The quotient is:", result)
	}
}
