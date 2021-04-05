package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Couple struct {
	Person1 Person
	Person2 Person
}

func (p Person) lastFirst() string {
	return p.LastName + ", " + p.FirstName
}

func (p Person) couple(other Person) Couple {
	return Couple{Person1: p, Person2: other}
}

func (c Couple) toString() string {
	return c.Person1.lastFirst() + " and " + c.Person2.lastFirst()
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
		fmt.Println(p2.lastFirst())
	}

	// coupling
	p1 := Person{FirstName: "Danielle", LastName: "Hollman"}
	p2 := Person{FirstName: "Greg", LastName: "Hollman"}
	c := p1.couple(p2)
	fmt.Println(c.toString())

}
