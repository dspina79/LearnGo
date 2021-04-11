package main

import "fmt"

type animal struct {
	name       string
	age        int
	animalType string
}

func newAnimal(name string, animalType string) *animal {
	a := animal{name: name}
	a.age = 0
	a.animalType = animalType

	return &a
}

func printAnimal(objAnimal animal) {
	fmt.Println("Name:", objAnimal.name, "\tType:", objAnimal.animalType, "\tAge:", objAnimal.age)
}

func main() {
	myAnimal := newAnimal("Fido", "Dog")
	printAnimal(*myAnimal)
}
