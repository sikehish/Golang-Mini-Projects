// PROJECT 8: Learning about maps
//MAPS: A key value based Data structure. They are dynamic DS
//VVI resource: https://www.golang-book.com/books/intro/6

package main

import "fmt"

func main() {
	//  websites:=map[string]string{} //-->map declaration
	// //OR
	// var websites map[string]string = map[string]string{}
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	//Accessing values
	fmt.Println(websites["Linkedin"] == "") //Basically it doesnt exist, hence empty string ouput
	fmt.Println(websites["Amazon Web Services"])

	//Adding a new key value pair
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println("After inserting linkedin: ", websites)

	//Deleting a key value pair
	delete(websites, "Google")
	fmt.Println("After deleting google: ", websites)

	//Maps vs structs: 2 major differences
	//1.A key in map can be of any data type, while in struct, its always a string
	//2. With structs, we have a predefined  DS, and we cant really delete/add a new key. Structs solve a different problem: to a have a collection of items, while maps store akey value pairs
	//Struct is more descriptive than maps

}
