// PROJECT 6: Project revolving around structs

package main

import (
	"fmt"

	"github.com/sikehish/Golang-Mini-Projects/Project6/user"
)

//VVI NOTE: FUNCTION VS METHODS
//FUNCTION(NOT METHOD): Accept a struct and operates on the struct
//METHOD: Uses a receiver(which is the struct on which the method has to be called)-->methods are defined on structs

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//Instantiating user struct

	// //Approach 1(You can also optionally omit fields, which would then be set to null)
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

	// //Approach 4
	// appUser := user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	//NOTE: There is another approach(using a function call like new to instantiate a struct)

	// var appUser *user
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// outputUserDetails(userFirstName, userLastName, userBirthdate)
	// outputUserDetails(appUser)
	// outputUserDetails(&appUser) //Pointer based approach
	//OR Method based approach
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	// // Instantiating after createing a seperate package for user
	// appUser2 := user.User{
	// 	firstName: userFirstName, --> firstName should have a capital 'F' in user struct for it to be accessible here
	// }

	//A new admin is created in user package
	appAdmin := user.NewAdmin("admin@gmail.com", "12345")

	// //Scenario 1: if User field is explicitly mentioned
	// appAdmin.User.OutputUserDetails()
	// appAdmin.User.ClearUserName()
	// appAdmin.User.OutputUserDetails()

	//Scenario 2: If User struct is embedded into the admin struct as it is
	appAdmin.OutputUserDetails()
	appAdmin.ClearUserName()
	appAdmin.OutputUserDetails()

}

//FUNCTIONS
// func outputUserDetails(firstName, lastName, birthdate string) {
// 	fmt.Println(firstName, lastName, birthdate)
// }
// Pass by value approach below
// func outputUserDetails(u user) {
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

// // Pass by ref approach below:
// func outputUserDetails(u *user.User) {
// 	//Pointers to struct exception: Normally you'd dereference to access the value like: (*u).birtthdate, which is valid, but in go, we have a shortcut/exception: (*u).birthdate is same as u.birthdate(As you can see in the below line)
// 	fmt.Println(u.firstName, u.lastName, u.birthdate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// fmt.Scan(&value) (doesnt take empty string as input, but Scanln stops recording the input as soon as new line is encountered)
	fmt.Scanln(&value)
	return value
}

// // -------
// // CUSTOM types OR(Type aliases)
// // NOTE YOU CANNOT ADD METHODS TO A PREDEFINED TYPE LIKE string or int

// type str string;

// func (text str) log(){
// 	fmt.Println(text)
// }

// func main(){
// 	var name str="Hisham"
// 	name.log()
// }
