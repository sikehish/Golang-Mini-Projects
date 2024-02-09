//Project 9:: Recursion

package main

import "fmt"

func main() {
	fact := factorial(5)
	factRec := factorialRecursive(5)
	fmt.Println(fact, factRec)
}

func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

func factorialRecursive(number int) int {
	if number == 1 {
		return 1
	}
	return number * factorialRecursive(number-1)
}
