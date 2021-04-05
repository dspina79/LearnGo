package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func getFullName(person Person) string {
	return person.FirstName + " " + person.LastName
}

func main() {
	p := Person{FirstName: "Dean", LastName: "Anips"}
	s := getFullName(p)
	fmt.Println(s)

	firstNames := []string{"Dean", "Mildred", "Elizabeth"}
	lastNames := []string{"Smith", "Oferi", "Gillian"}
	persons := []Person{}
	for _, fn := range firstNames {
		for _, ln := range lastNames {
			p1 := Person{FirstName: fn, LastName: ln}
			persons = append(persons, p1)
		}
	}

	for _, p2 := range persons {
		s2 := getFullName(p2)
		fmt.Println(s2)
	}

}
