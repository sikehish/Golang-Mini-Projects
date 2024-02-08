// PROJECT 8: Learning about arrays,slices and maps
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

}
