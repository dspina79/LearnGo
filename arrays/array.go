package main

import "fmt"

func main() {
	var firstArray [5]int
	for i := 0; i < len(firstArray); i++ {
		firstArray[i] = i * 10
	}

	fmt.Println(firstArray)

	secondArray := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(firstArray); i++ {
		secondArray[i] = secondArray[i] * 5
	}

	fmt.Println(secondArray)

	var twoDimArray [12][12]int

	for x := 0; x < 12; x++ {
		for y := 0; y < 12; y++ {
			twoDimArray[x][y] = (x + 1) * (y + 1)
		}
	}

	fmt.Println(twoDimArray)
}
