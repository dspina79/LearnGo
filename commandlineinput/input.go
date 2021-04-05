package main

import "fmt"

func double(x *int) {
	*x *= 2
}

func removeDecimal(x *float32) {
	tempX := int(*x)
	*x = float32(tempX)
}

func main() {
	var name string
	var age int
	var decVal float32
	fmt.Print("Enter in your name: ")
	fmt.Scan(&name)
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("Hello", name, ". Your age is", age)
	double(&age)
	fmt.Println("The computer has doubled your age. It is now", age)
	fmt.Print("Enter a decimal value: ")
	fmt.Scan(&decVal)
	fmt.Println("Current value of the decimal is ", decVal)
	removeDecimal(&decVal)
	fmt.Println("Removal of the decimal makes it ", decVal)
}
