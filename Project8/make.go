//PROJECT8: Learning about make function and type aliases

package main

import "fmt"

func main() {

	//How we've created slices so far:
	userNames := []string{} //no values initialized here but that can be done
	//Behind the scenes, in order to manage the slices, go creates arrays. SO the underlying DS used in slices is array
	userNames = append(userNames, "Marlon")
	userNames = append(userNames, "Simeone")
	fmt.Println(userNames)

	//Using make: make is used to preallocate some space , used in scenarios where we are aware of the number of elements ythat will be inserted, and thus the expansion of the slices takes place in an efficient manner as the slice need not be expanded at each stage like in the case []string{} where go has no idea about the length to be intialized, and this could lead to inefficiencies. In the older approach, if the array pace is exhausted, go would create a new array having a greater size, but using make, we can avoid this, atleast at the vbeginning of data insertion

	userNames = make([]string, 2) //initial length of the slice set to 2 and intialised to empty strings. Thus we need not append as two slots are already present , and we shall only append if 3rd element is to be inserted.If we append the first 2 elements, then we'll have 4 elements, with the first 2 elements being empty strings
	userNames[0] = "Julie"
	userNames[1] = "Sanjay"
	fmt.Println(userNames)

	//More efficient usage of make: Making use of a third param: capacity
	userNames = make([]string, 2, 5)
	fmt.Println(userNames, len(userNames)) //Just like the above make function, 2 empty slots are intialised.Initially, the slice has space for two elements.underlying array that backs the slice has a capacity of five elements. The capacity is the maximum number of elements that the slice can hold without allocating additional memory.In summary, userNames is a slice of strings with an initial length of 2 and a capacity of 5. The underlying array is created with a capacity of 5, but only the first 2 elements are considered part of the slice initially. If you were to add more elements to the slice, it could grow up to the specified capacity before a new underlying array is allocated.
	userNames[0] = "Julie"
	userNames[1] = "Sanjay"
	// userNames[2] = "Krishna" //->This is still not possible
	userNames = append(userNames, "Krishna")
	fmt.Println(userNames)

	//Creating maps using make
	courseRatings := make(map[string]float64)
	//In this case, Go will have to reallocate memory whenever we add new items to the map, thus there would be no difference between make approach and map[string]float64{}
	courseRatings["CSS"] = 4.6
	courseRatings["Go"] = 4.9
	fmt.Println(courseRatings)

	//Better approach: Using make for efficent mem allocation
	courseRatings = make(map[string]float64, 3) //3 is the initial capacity of the map. It means that the map is initially allocated to have space for three key-value pairs. The capacity is the number of key-value pairs that the map can hold without resizing.In summary, courseRatings is a map with string keys and float64 values, and it is initially allocated with a capacity of 3. The initial capacity is not a strict limit on the number of elements the map can hold; it just provides an initial size to avoid immediate resizing in case you know an approximate size for your map. The map can dynamically grow beyond the initial capacity as needed.
	courseRatings["CSS"] = 4.6
	courseRatings["Go"] = 4.9
	courseRatings["React"] = 4.4
	//AFter the third key value pair is inserted, and when we insert the below key value pair, go will reallocate memory for the map
	courseRatings["Express"] = 4.8

}
