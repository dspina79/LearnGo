package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Dave"] = 101
	m["Dean"] = 102

	fmt.Println(m)

	fmt.Println("Changing Dean to 201")
	m["Dean"] = 201
	fmt.Println(m)
	fmt.Println("Adding new people")
	m["Sharon"] = 301
	m["Peggy"] = 422
	fmt.Println(m)

	fmt.Println("Removing Sharon")
	delete(m, "Sharon")
	fmt.Println(m)

	fmt.Println("Accessing invalid entry")
	_, something := m["something"]
	fmt.Println("Something =", something)

	if something {
		fmt.Println("Something was found!")
	} else {
		fmt.Println("Something was not found! It doesn't exist in the map.")
	}

}
