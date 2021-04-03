package main

import "fmt"

func main() {
	s := "About to do for loops" // initializes variable without the need for var x
	fmt.Println(s)
	for i := 0; i < 3; i++ {
		fmt.Println("The index is at ", i)
	}
	fmt.Println("Printing even numbers 1 through 20")
	for j := 1; j <= 20; j++ {
		if j%2 == 0 {
			fmt.Println(j)
		}
	}
}
