//A seperate package for user struct

package user

import (
	"errors"
	"fmt"
	"time"
)

// can be imported directly
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Struct embedding: Using a struct in another struct

// 1. Anonymous Embedding
type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// // OR 2. Explicit embedding
// type Admin struct {
// 	email    string
// 	password string
// 	User     User
// }

// func NewAdmin(email, password string) Admin {
// 	return Admin{
// 		email:    email,
// 		password: password,
// 		User: User{
// 			firstName: "ADMIN",
// 			lastName:  "ADMIN",
// 			birthdate: "---",
// 			createdAt: time.Now(),
// 		},
// 	}
// }

// NOTE: methods can be defined on both value receivers (non-pointer receivers) and pointer receivers. When you define a method on a type, you can use either a value receiver or a pointer receiver. Go automatically dreferences/creates a pointer to match the receicer's type. So, we need not really worry about pointers vs values in Go structs

// Method (Value receiver)
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

// Mutation methods(they dont create a copy but instead take a pointer and modify the data)(Pointer receiver)
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// Constructor function is a utility function which starts with "New" (or "new") which is just a convention, tpo instantite a struct
//
//	func newUser(firstName, lastName, birthdate string) user {
//		return user{
//			firstName: firstName,
//			lastName:  lastName,
//			birthdate: birthdate,
//			createdAt: time.Now(),
//		}
//	}
//
// PASS BY REFERENCE VERSION OF THE ABOVE CONSTRUCTOR(Not required tho)
// NOTE: NewUser or New--> it doesnt matter--> New is a very common convention used to name constrcutor functions,like errors.New()
// func NewUser(firstName, lastName, birthdate string) (*User, error)  OR
func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("No field can be left empty")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

//NOTE: User{} creates an empty struct
