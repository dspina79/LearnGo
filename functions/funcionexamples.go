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

func main() {
	var x int = 10
	var y int = 15

	fmt.Println(x, " + ", y, "=", add(x, y))
	fmt.Println(x, " x ", y, "=", multiply(x, y))
	fmt.Println(x, " / ", y, "=", divide(x, y))

	values := []int{23, 4, 234, 234, 234, 45, -2, 3452, 532, 901, 32, 12, 31}
	fmt.Println("Array is: ", values)
	min, max := getMinMax(values)

	fmt.Println("The min of hte array is: ", min)
	fmt.Println("The max of hte array is: ", max)

}
