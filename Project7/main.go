//Project 7: Building a notes application(CLI) with the help of structs

package main

import (
	"fmt"

	"github.com/sikehish/Golang-Mini-Projects/Project7/note"
)

func main() {
	// var title, content, err = getNoteData()
	//OR //  title, content, err := getNoteData()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(title, content)

	title, content := getNoteData() //OR var title,content=func() OR var title, content string = func()

	newNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newNote)

}

// func getNoteData() (string, string, error) {
// 	title, err := getUserInput("Note title: ")
// 	if err != nil {
// 		return "", "", err
// 	}
// 	content, err := getUserInput("Note content: ")
// 	if err != nil {
// 		return title, "", err
// 	}

// 	return title, content, nil
// }

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

// func getUserInput(prompt string) (string, error) {
// 	var input string
// 	fmt.Print(prompt)
// 	fmt.Scanln(&input)

// 	if input == "" {
// 		return "", errors.New("Invalid input")
// 	}
// 	return input, nil
// }

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
