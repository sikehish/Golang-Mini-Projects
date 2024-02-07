//Project 7: Building a notes application(CLI) with the help of structs

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

	newNote.Display()

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

	// fmt.Scanln(&input)
	// //We get unexpected behaviour when we enter "Learn Go" as input for the Note title prompt
	// //This is because all the scan functions are designed to accept signle word input and beyond that(space seperated words) we might not get the desired functionality

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSpace(text) //Removes leading and tralinign whitespaces
	// //OR
	// input = strings.TrimSuffix(text, "\n")
	// input = strings.TrimSuffix(input, "\r") //For windows OS

	return input
}
