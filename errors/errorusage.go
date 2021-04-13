package main

import (
	"errors"
	"fmt"
)

type myError struct {
	errorNum     int
	errorMessage string
}

func (myErr *myError) Error() string {
	return fmt.Sprintf("%d: %s", myErr.errorNum, myErr.errorMessage)
}

func divde(x, y int) (float32, error) {
	if y == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return float32(x) / float32(y), nil
}

func checkPassword(pass string) (bool, error) {
	if pass == "" {
		// can't allow blank passwords
		return false, &myError{errorNum: 0, errorMessage: "Password cannot be blank."}
	}

	if pass == "correctpass" {
		return true, nil
	} else {
		return false, &myError{errorNum: 1, errorMessage: "Password is incorrect."}
	}

}

func testPassword() {
	var password string

	fmt.Print("Enter the password: ")
	fmt.Scanf("%s", password)

	passResult, passError := checkPassword(password)
	if passResult {
		fmt.Println("Password is correct.")
	} else {
		fmt.Println(passError)
	}
}

func main() {
	var x int
	var y int

	fmt.Print("Enter a numerator: ")
	fmt.Scanf("%d", &x)
	fmt.Print("Enter a denominator: ")
	fmt.Scanf("%d", &y)

	result, err := divde(x, y)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The quotient is:", result)
	}

}
