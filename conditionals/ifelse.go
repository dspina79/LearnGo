package main

import "fmt"

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
}
