package main

import "fmt"

type animal struct {
	name       string
	age        int
	animalType string
}

func (a *animal) incrementAge() {
	a.age += 1
}

func (a animal) getNameType() string {
	return a.name + "\tType: " + a.animalType
}

func (a animal) toString() string {
	return "Name: " + a.name + "\tType: " + a.animalType + "\tAge: " + fmt.Sprint(a.age)
}

func newAnimal(name string, animalType string) *animal {
	a := animal{name: name}
	a.age = 0
	a.animalType = animalType

	return &a
}

func printAnimal(objAnimal animal) {
	fmt.Println(objAnimal.toString())
}

func main() {
	myAnimal := newAnimal("Fido", "Dog")
	printAnimal(*myAnimal)

	mySecondAnimal := animal{name: "Gerald", animalType: "Cat", age: 12}
	myThirdAnimal := animal{"Minne", 8, "Leopard"} // follows the order of the structs
	printAnimal(mySecondAnimal)
	printAnimal(myThirdAnimal)

	fmt.Println("Incrememnt all ages")
	myAnimal.incrementAge()
	mySecondAnimal.incrementAge()
	myThirdAnimal.incrementAge()

	printAnimal(*myAnimal)
	printAnimal(mySecondAnimal)
	printAnimal(myThirdAnimal)

}
