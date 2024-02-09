// Project 9: Learn about Functions
// //NOTE: Go is a programming language that supports first-class functions. This means that functions in Go can be:

//     Assigned to variables.
//     Passed as arguments to other functions.
//     Returned as values from other functions.

package main

import "fmt"

// Type aliases
type transform func(int) int

// Another example:(We're not using it but just an example)
type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	// doubled := doubleNumbers(numbers)
	// fmt.Println(doubled)

	//Functions can be passed as arguments
	doubled := transformNumbers(numbers, double)
	tripled := transformNumbers(numbers, triple)
	fmt.Println(doubled, tripled)

	// Functions in go can also be returned:
	transformerFn1 := getTransformerFunc(numbers)
	transformerFn2 := getTransformerFunc([]int{2, 3, 5})
	transformedNumbers := transformNumbers(numbers, transformerFn1)
	moreTransformedNumbers := transformNumbers([]int{2, 3, 5}, transformerFn2)

	fmt.Println(transformedNumbers, moreTransformedNumbers)
}

//	func doubleNumbers(numbers []int) []int {
//		// var dNumbers []int
//		//OR
//		dNumbers := []int{}
//		for _, val := range numbers {
//			dNumbers = append(dNumbers, double(val))
//		}
//		return dNumbers
//	}
//
// OR
//
//	func transformNumbers(numbers []int, transformFunc func(int) int) []int {
//		dNumbers := []int{}
//		for _, val := range numbers {
//			dNumbers = append(dNumbers, transformFunc(val))
//		}
//		return dNumbers
//	}
//
// OR
func transformNumbers(numbers []int, transformFunc transform) []int {
	dNumbers := []int{}
	for _, val := range numbers {
		dNumbers = append(dNumbers, transformFunc(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// func getTransformerFunc(numbers []int) func(int) int{}
// OR
func getTransformerFunc(numbers []int) transform {
	if numbers[0] == 1 {
		return double
	} else {
		return triple
	}
}
