package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Dave"

	if name == "Dave" {
		fmt.Println("Hello,", name)
	} else if name == "Dean" {
		fmt.Println("Oh no, it's", name)
	} else {
		fmt.Println(name, "is a nice name.")
	}

	// allow statements before the num
	if lastName := "Smith"; name == "Ferguson" {
		fmt.Println("You're one of the", lastName, "s")
	} else {
		fmt.Println("So, you're", name, lastName)
	}

	// switch
	var b int = 3

	switch b {
	case 1:
		fmt.Println("The number one was selected.")
	case 2:
		fmt.Println("The number two was selected.")
	case 3:
		fmt.Println("The number three was selected.")
	default:
		fmt.Println("Some other number was selected.")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.")
	default:
		fmt.Println("It's part of the work week.")
	}

	// delegated function
	add := func(x int, y int) int {
		return x + y
	}

	result := add(3, 5)
	fmt.Println("The result is ", result)
}
