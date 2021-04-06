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

}
