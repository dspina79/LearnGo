package main

import "fmt"

type threeDimensionalObject interface {
	volume() float32
}

type cube struct {
	side int
}

type box struct {
	length, width, height int
}

type sphere struct {
	radius float32
}

func (s sphere) volume() float32 {
	return (4.0 / 3.0) * 3.14159 * (s.radius * s.radius * s.radius)
}

func (b box) volume() float32 {
	return float32(b.height * b.length * b.width)
}

func (c cube) volume() float32 {
	return float32(c.side * c.side * c.side)
}

func printVolume(obj threeDimensionalObject, name string) {
	fmt.Println("The of volume of the", name, "is", obj.volume(), "cubic units")
}

func main() {
	s := sphere{radius: 4}
	c := cube{side: 3}
	b := box{3, 4, 5}

	printVolume(s, "SPHERE")
	printVolume(c, "CUBE")
	printVolume(b, "BOX")
}
