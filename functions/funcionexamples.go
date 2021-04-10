package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) float32 {
	fx := float32(x)
	fy := float32(y)

	return fx / fy
}

func addGroup(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func getMinMax(vals []int) (int, int) {
	min := 0
	max := 0

	for i, val := range vals {
		if i == 0 || val < min {
			min = val
		}
		if i == 0 || val > max {
			max = val
		}
	}

	return min, max
}

func averageGroup(nums ...int) float32 {
	sum := addGroup(nums...)
	var a float32 = float32(sum) / float32(len(nums))

	return a
}

// recursion

func fact(x int) int {
	if x < 2 {
		return 1
	} else {
		return x * fact(x-1)
	}
}

func main() {
	var x int = 10
	var y int = 15

	fmt.Println(x, " + ", y, "=", add(x, y))
	fmt.Println(x, " x ", y, "=", multiply(x, y))
	fmt.Println(x, " / ", y, "=", divide(x, y))

	values := []int{23, 4, 234, 234, 234, 45, -2, 3452, 532, 901, 32, 12, 31}
	fmt.Println("Array is: ", values)
	min, max := getMinMax(values)

	fmt.Println("The min of the array is: ", min)
	fmt.Println("The max of the array is: ", max)

	sum := addGroup(2, 3, 4, 5, 6, 7, 10, 12, 13)
	average := averageGroup(2, 3, 4, 5, 6, 7, 10, 12, 13)
	fmt.Println("The sum of 2,3,4,5,6,7,10,12, and 13 is ", sum)
	fmt.Println("The average of 2,3,4,5,6,7,10,12, and 13 is ", average)

	base := 6
	fact6 := fact(base)

	fmt.Println("The factorial of", base, "is", fact6)

}
