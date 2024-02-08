package main

import "fmt"

//Solution: https://github.com/mschwarzmueller/go-complete-guide-resources/blob/main/code/07-arrays-slices-maps/07-exercise-solution/practice.go

// Task 7(struct)
type Product struct {
	title string
	id    string
	price float64
}

func main() {

	//Task 1
	hobbies := [3]string{"Coding", "Photography", "Football"} //OR var hobbies = [3]string{"Coding", "Photography", "Football"} OR var hobbies [3]string = [3]string{"Coding", "Photography", "Football"}
	fmt.Println("My Hobbies: ", hobbies)
	fmt.Println("----------------------------------------------------------")

	//Task 2
	ele := hobbies[0:1] //OR hobbies[:1]//OR, if you dont want a slice, and only want the element, ele := hobbies[0]
	arr := hobbies[1:]  //OR arr:=hobbies[1:3]
	fmt.Printf("%v\n%v\n", ele, arr)
	fmt.Println("----------------------------------------------------------")

	//Task 3
	slc1 := ele[:2]
	slc2 := ele[0:2]
	fmt.Println(slc1, slc2)
	fmt.Println("----------------------------------------------------------")

	//Task 4
	slc := slc1[1:3]
	fmt.Println(slc)
	fmt.Println("----------------------------------------------------------")

	//Task 5
	courseGoals := []string{"Get a good score", "Get a good job"}
	fmt.Println(courseGoals)
	fmt.Println("----------------------------------------------------------")

	//Task 6
	courseGoals[1] = "Get a high paying job"
	courseGoals = append(courseGoals, "Make good friends") //Overwrites the exisitng slice
	// courseGoals2 := append(courseGoals, "Make good friends") //OR we can store the result in a different slice, as append doesnt modify the exisiting slice, insted it returns a new one
	fmt.Println(courseGoals)
	fmt.Println("----------------------------------------------------------")

	//Task 7(Bonus)
	productList := []Product{Product{"Harpic", "ABC123", 180.0}, Product{"Dettol", "ABC256", 185.5}}
	// OR productList := []Product{{"Harpic", "ABC123", 180.0}, {"Dettol", "ABC256", 185.5}}
	fmt.Println(productList)
	productList = append(productList, Product{"Garlic", "DEF234", 25.74})
	fmt.Println(productList)
	fmt.Println("----------------------------------------------------------")
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
