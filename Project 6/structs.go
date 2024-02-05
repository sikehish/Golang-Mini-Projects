// PROJECT 6: Project revolving around structs

package main

import (
	"fmt"
	"time"
)

//VVI NOTE: FUNCTION VS METHODS
//FUNCTION(NOT METHOD): Accept a struct and operates on the struct
//METHOD: Uses a receiver(which is the struct on which the method has to be called)-->methods are defined on structs

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

//Method
func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//Instantiating user struct

	// //Approach 1(You can also optionally omiy fields, which would then be set to null)
	// var appUser user
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// OR
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// //Approach 2
	// var appUser user = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	// // Approach 3
	// var appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	//Approach 4
	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	//NOTE: There is another approach(using a function call like new to instantiate a struct)

	// outputUserDetails(userFirstName, userLastName, userBirthdate)
	// outputUserDetails(appUser)
	// outputUserDetails(&appUser) //Pointer based approach
	//OR Method based approach
	appUser.outputUserDetails()

}

//FUNCTIONS
// func outputUserDetails(firstName, lastName, birthdate string) {
// 	fmt.Println(firstName, lastName, birthdate)
// }
// Pass by value approach below
// func outputUserDetails(u user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

//Pass by ref approach below:
func outputUserDetails(u *user) {
	//Pointers to struct exception: Normally you'd dereference to access the value like: (*u).birtthdate, which is valid, but in go, we have a shortcut/exception: (*u).birthdate is same as u.birthdate(As you can see in the below line)
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
