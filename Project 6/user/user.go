//A seperate package for user struct

package user

import (
	"errors"
	"fmt"
	"time"
)

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

//Mutation methods(they dont create a copy but instead take a pointer and modify the data)
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

//Constructor function is a utility function which starts with "New" (or "new") which is just a convention, tpo instantite a struct
// func newUser(firstName, lastName, birthdate string) user {
// 	return user{
// 		firstName: firstName,
// 		lastName:  lastName,
// 		birthdate: birthdate,
// 		createdAt: time.Now(),
// 	}
// }
//PASS BY REFERENCE VERSION OF THE ABOVE CONSTRUCTOR(Not required tho)
func NewUser(firstName, lastName, birthdate string) (*user, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("No field can be left empty")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}