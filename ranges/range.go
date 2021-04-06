package main

import "fmt"

func main() {
	sum := 0
	nums := []int{1, 2, 3, 4, 5, 6, 7, 15, 201, 329}

	for i, num := range nums {
		fmt.Println("The number at index", i, "is", num)
	}

	for _, num := range nums {
		sum += num
	}

	fmt.Println("The sum is", sum)
	average := float32(sum) / float32(len(nums))
	fmt.Println("The average is", average)

	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	m["Z"] = 26

	fmt.Println("The map contents are", m)

	for k, v := range m {
		fmt.Println("The key", k, "has the value", v)
	}

}
