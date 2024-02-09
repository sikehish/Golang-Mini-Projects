// Project 9: Anonymous functions, Closures
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	//Exmaple of anonymous Functions(has no name and is defined wherever needed)
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformed)

	//Closures
	double := createTransformer(2)
	triple := createTransformer(3)

	doubledArr := transformNumbers(&numbers, double)
	tripledArr := transformNumbers(&numbers, triple)
	fmt.Println(doubledArr, tripledArr)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// Closures
// createTransformer is a factory function.In Go, a factory function is a function that returns a new instance of a type. The purpose of a factory function is to encapsulate the process of creating an object, providing flexibility and abstraction for the user.For example, the New() function that we used with structs are factory funcs
// Functions are closures in go.closure is a function value that references variables from its containing function's body. In Go, if you define a function inside another function and it references variables from the outer function, it becomes a closure.
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
