// Project 9:... operator,variadics, and unpacking(or spread operator in JS)operator in Go

package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	sum := sumup(numbers)
	fmt.Println(sum)

	//Using variadic functions
	sum = sumup1(1, 10, 15)
	fmt.Println(sum)

	sum = sumup2(1, 10, 15)
	fmt.Println(sum)

	//Splitting slices into parameter values
	sum = sumup1(numbers...)
	fmt.Println(sum)

	sum = sumup2(1, numbers...)
	fmt.Println(sum)
}

func sumup(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

// Variadic functions
func sumup1(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func sumup2(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, val := range numbers {
		sum += val
	}
	return sum
}
