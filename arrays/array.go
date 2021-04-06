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

	for a := 0; a < 12; a++ {
		for b := 0; b < 12; b++ {
			fmt.Print(twoDimArray[a][b], "\t")
		}
		fmt.Print("\n")
	}

	// slices
	slice1 := make([]string, 10)
	slice1[0] = "Mark"
	slice1[1] = "Rick"
	slice1[2] = "Steve"
	slice1[3] = "Lynn"
	slice1[4] = "Sharon"
	slice1[5] = "Esther"
	slice1[6] = "Ryan"
	slice1[7] = "Bea"
	slice1[8] = "Milton"
	slice1[9] = "Cheryl"

	fmt.Println(slice1)

	fmt.Println("Slice [2:5] = ", slice1[2:5])
	fmt.Println("Slice [5:] = ", slice1[5:])
	fmt.Println("Slice [9:] = ", slice1[9:])
	fmt.Println("Slice [:4] = ", slice1[:4])

}
