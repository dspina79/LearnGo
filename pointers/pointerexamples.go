package main

import "fmt"

func dbl(x *int) {
	*x *= 2
}

func getGrainWeight(numGrains int) float32 {
	weightGrainKg := (1.0 / 64.0) / 1000.00
	weightGransLbs := weightGrainKg * 2.2
	return float32(numGrains) * float32(weightGransLbs)
}

func main() {
	num := 1

	fmt.Println("We are starting with", num)

	for i := 0; i < 64; i++ {
		dbl(&num)
		fmt.Println("Double that is", num, "Which, in grains of rice = ", getGrainWeight(num), "pounds.")
	}

}
