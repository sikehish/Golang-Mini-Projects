// PROJECT 8: Learning about arrays and slices
package main

import "fmt"

func main() {
	//Arrays
	// var productName [4]string //4 empty slots/strings
	// productName = [4]string{"A", "B", "C", "D"}
	//OR
	var products [4]string = [4]string{"A", "B", "C", "D"} //OR
	var productName [4]string = [4]string{"A Book"}        //OR Initialising the array with only 1 value, and the remaining slots are initialised with empty string
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(products)
	fmt.Println(productName)

	// In an array literal, the ... notation specifies a length equal to the number of elements in the literal.
	prices = [...]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println("Array Literal: ", prices)

	//Creating a slice from the array
	// featuredPrices := prices[1:3] //Includes elements starting from index 1 to index 3-1=2.
	// featuredPrices := prices[:3]   //From index 0 to 2
	featuredPrices := prices[1:] //From index 1 to las idx=3
	//NOTE: negative indices and indices larger than the array size arent allowed
	fmt.Println(featuredPrices, len(featuredPrices), cap(featuredPrices))

	//Creating a slice from a slice:
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices, len(highlightedPrices), cap(highlightedPrices))
	//In the above example, len=1 and capacity=3. This is because the slice from which it is sliced has 3 elements, and here were obtaining just the first element. We can also extend highlightedPrices towards the right and can span across all 3 elements instead of just 1. Note that if highlighted prices would sotre the slice from 1:2(the featured Prices has 3 elements), then we would only get access to the 1st element, with an option to extend/span across index 2 as well starting from index 1, but we cannot span index 0 element as it is towards the left of the starting index. Below are a few exampels:

	//As capacity of highlightedPrices in the above example is 3,we can do the following
	highlightedPrices = highlightedPrices[0:] //we obtin just the first element, but if we explictly mention the ending index, and if its less than or equal to capacity, then we shall see other elements stored in the base slice
	highlightedPrices = highlightedPrices[0:3]
	fmt.Println(featuredPrices, highlightedPrices, len(highlightedPrices), cap(highlightedPrices))

	highlightedPrices = featuredPrices[1:2]
	fmt.Println(featuredPrices, highlightedPrices, len(highlightedPrices), cap(highlightedPrices))
	//Now, highlightedPrices:=highlightedPrices[0:2] will fetch index 1 an index 2 elements of the featured prices, and there's no way that we'd be able to access any elements to the left of index1 as slice starts from index 1

	//Also, slices are mutable references/windows of the original array. They arent copies but are references. THus modifying a slice would resultin modification of the base slice(if it exists and it does in the below case(featuredPrices)) and base array. Ex:
	highlightedPrices[0] = 1000
	fmt.Printf("\n\nResult of Mutation:\nPrices: %v\nfeaturedPrices:%v\nhighlightedPrices:%v\n", prices, featuredPrices, highlightedPrices)

	//SLICES(They are dynamic arrays with no predefined size)
	priceSlice := []float64{10.99, 8.99}
	fmt.Println(priceSlice[0:1])
	fmt.Println(priceSlice[0])
	priceSlice[1] = 9.99
	// priceSlice[2]=100 --> invalid as only 2 elements are present. If you want to append a new value, then use append()
	updatedPrices := append(priceSlice, 5.99) //returns a brand new slice and doesnt modify the underlying slice.Also we can append n number of values, that is:
	updatedPrices = append(priceSlice, 65, 4545, 32)
	fmt.Println(priceSlice, updatedPrices)
	//You can modify the existing slice by overwriting it: priceSlice := append(priceSlice, 5.99)

	//NOTE: Slices are more widely used than arrays, given the flexibility

	// Removing en element from a slice
	fmt.Printf("\nRemoving an element in a slice:\nOriginal Slice: %v\n", updatedPrices)
	updatedPrices = updatedPrices[1:] //-->removes the first element
	fmt.Printf("Slice after eliminating the first element: %v\n", updatedPrices)
	updatedPrices = updatedPrices[:len(updatedPrices)-1] //-->removes the first element
	fmt.Printf("Slice after eliminating the last element and the first element: %v\n", updatedPrices)

	//To create a copy of a slice in Go, you can use the copy function or simply re-slice the original slice. The copy function is often preferred when you want to ensure that the original and copied slices are completely independent of each other: https://freshman.tech/snippets/go/copy-slices/

	//Unpacking list values
	//Appending one slice to another slice
	discountPrices := []float64{1.23, 12.3, 123}
	//Note: append only works on slices and not on arrays
	appendedSlice := append(updatedPrices, discountPrices...) //https://yourbasic.org/golang/three-dots-ellipsis/
	//... is the unpacking operator(known as spread in JS) seperates the elements of the slice (comma seperated), and hence can be used as arguments for the append() func
	fmt.Println(updatedPrices, appendedSlice)
}
